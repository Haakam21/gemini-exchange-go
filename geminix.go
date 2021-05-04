package geminix

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BaseUrl          = "https://api.gemini.com"
	WsBaseUrl        = "wss://api.gemini.com"
	SandboxBaseUrl   = "https://api.sandbox.gemini.com"
	SandboxWsBaseUrl = "wss://api.sandbox.gemini.com"

	// public
	SymbolsUri        = "/v1/symbols"
	SymbolDetailsUri  = "/v1/symbols/details"
	TickerUri         = "/v2/ticker/%s"
	CandlesUri        = "/v2/candles/%s/%s"
	OrderBookUri      = "/v1/book/%s"
	TradesUri         = "/v1/trades/%s"
	AuctionUri        = "/v1/auction/%s"
	AuctionHistoryUri = "/v1/auction/%s/history"

	// roles
	RolesUri = "/v1/roles"

	// order placement
	NewOrderUri      = "/v1/order/new"
	CancelOrderUri   = "/v1/order/cancel"
	CancelSessionUri = "/v1/order/cancel/session"
	CancelAllUri     = "/v1/order/cancel/all"

	// order status
	OrderStatusUri  = "/v1/order/status"
	ActiveOrdersUri = "/v1/orders"
	PastTradesUri   = "/v1/mytrades"

	// fee and volume
	NotationalVolumeUri = "/v1/notationalvolume"
	TradeVolumeUri      = "/v1/tradevolume"

	// clearing
	NewClearingOrderUri     = "/v1/clearing/new"
	NewBrokerOrderUri       = "/v1/broker/new"
	ClearingOrderStatusUri  = "/v1/clearing/status"
	CancelClearingOrderUri  = "/v1/clearing/cancel"
	ConfirmClearingOrderUri = "/v1/clearing/confirm"

	// fund management
	BalancesUri           = "/v1/balances"
	NotationalBalancesUri = "/v1/notionalbalances/%s"
	TransfersUri          = "/v1/transfers"
	DepositAddressesUri   = "/v1/addresses/%s"
	NewDepositAddressUri  = "/v1/deposit/%s/newAddress"
	WithdrawCryptoUri     = "/v1/withdraw/%s"
	InternalTransferUri   = "/v1/account/transfer/%s"
	AddBankUri            = "/v1/payments/addbank"
	PaymentMethodsUri     = "/v1/payments/methods"

	// approved address
	RequestAddressUri    = "/v1/approvedAddresses/%s/request"
	ApprovedAddressesUri = "/v1/approvedAddresses/account/%s"
	RemoveAddressUri     = "/v1/approvedAddresses/%s/remove"

	// account administration
	AccountDetailUri = "/v1/account"
	CreateAccountUri = "/v1/account/create"
	AccountsUri      = "/v1/account/list"

	// session
	HeartbeatUri = "/v1/heartbeat"
)

type Network string

const (
	Bitcoin     Network = "bitcoin"
	Ethereum    Network = "ethereum"
	BitcoinCash Network = "bitcoincash"
	Litecoin    Network = "litecoin"
	ZCash       Network = "zcash"
	Filecoin    Network = "filecoin"
)

type Client struct {
	url    string
	key    string
	secret string
}

func NewClient(key string, secret string, sandbox bool) *Client {
	var url string
	if sandbox {
		url = SandboxBaseUrl
	} else {
		url = BaseUrl
	}

	return &Client{url: url, key: key, secret: secret}
}

// nonce returns a generic nonce based on unix timestamp
func nonce() int64 {
	return time.Now().UnixNano()
}

type ApiError struct {
	Reason  string
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("[%v] %v", e.Reason, e.Message)
}

type Response struct {
	Result string
	ApiError
}

// BuildHeader handles the conversion of post parameters into headers formatted
// according to Gemini specification. Resulting headers include the API key,
// the payload and the signature.
func (c *Client) BuildHeader(req *map[string]interface{}) (http.Header, error) {

	reqStr, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	payload := base64.StdEncoding.EncodeToString([]byte(reqStr))

	mac := hmac.New(sha512.New384, []byte(c.secret))
	mac.Write([]byte(payload))

	signature := hex.EncodeToString(mac.Sum(nil))

	header := http.Header{}
	header.Set("X-GEMINI-APIKEY", c.key)
	header.Set("X-GEMINI-PAYLOAD", payload)
	header.Set("X-GEMINI-SIGNATURE", signature)

	return header, nil
}

// request makes the HTTP request to Gemini and handles any returned errors
func (c *Client) request(verb string, uri string, params map[string]interface{}) ([]byte, error) {
	url := c.url + uri

	req, err := http.NewRequest(verb, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	if params != nil {
		if verb == "GET" {
			q := req.URL.Query()
			for key, val := range params {
				q.Add(key, val.(string))
			}
			req.URL.RawQuery = q.Encode()
		} else {
			req.Header, err = c.BuildHeader(&params)
			if err != nil {
				return nil, err
			}
		}
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	json.Unmarshal(body, &res)
	if res.Result == "error" {
		return nil, &res.ApiError
	}

	return body, nil
}

func (c *Client) PublicRequest(uri string) ([]byte, error) {
	body, err := c.request("GET", uri, nil)

	return body, err
}

func (c *Client) PrivateRequest(uri string, params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = map[string]interface{}{
			"request": uri,
			"nonce":   nonce(),
		}
	} else {
		params["request"] = uri
		params["nonce"] = nonce()
	}

	body, err := c.request("POST", uri, params)
	return body, err
}

type Id string

// Id has a custom Unmarshal since it needs to handle unmarshalling from both
// string and int json types. This package takes the position that throughout
// ids should be strings and converted from json into strings where needed.
func (id *Id) UnmarshalJSON(b []byte) error {

	if len(b) > 0 && b[0] == '"' {
		b = b[1:]
	}

	l := len(b)
	if l > 0 && b[l-1] == '"' {
		b = b[:l-1]
	}

	*id = Id(b)
	return nil
}

type Order struct {
	OrderId           Id      `json:"order_id"`
	ClientOrderId     string  `json:"client_order_id"`
	Symbol            string  `json:"symbol"`
	Side              string  `json:"side"`
	Type              string  `json:"type"`
	Timestamp         int64   `json:"timestampms"`
	IsLive            bool    `json:"is_live"`
	IsCancelled       bool    `json:"is_cancelled"`
	IsHidden          bool    `json:"is_hidden"`
	WasForced         bool    `json:"was_forced"`
	Price             float64 `json:"price,string"`
	ExecutedAmount    float64 `json:"executed_amount,string"`
	RemainingAmount   float64 `json:"remaining_amount,string"`
	OriginalAmount    float64 `json:"original_amount,string"`
	AvgExecutionPrice float64 `json:"avg_execution_price,string"`
}

type Trade struct {
	OrderId       Id      `json:"order_id"`
	TradeId       Id      `json:"tid"`
	Timestamp     int64   `json:"timestampms"`
	Exchange      string  `json:"exchange"`
	Type          string  `json:"type"`
	FeeCurrency   string  `json:"fee_currency"`
	FeeAmount     float64 `json:"fee_amount,string"`
	Amount        float64 `json:"amount,string"`
	Price         float64 `json:"price,string"`
	IsAuctionFill bool    `json:"is_auction_fill"`
	Aggressor     bool    `json:"aggressor"`
	Broken        bool    `json:"broken"`
	Break         string  `json:"break"`
}

type Ticker struct {
	Bid    float64      `json:"bid,string"`
	Ask    float64      `json:"ask,string"`
	Last   float64      `json:"last,string"`
	Volume TickerVolume `json:"volume"`
}

type TickerVolume struct {
	BTC       float64 `json:",string"`
	ETH       float64 `json:",string"`
	USD       float64 `json:",string"`
	Timestamp int64   `json:"timestamp"`
}

type TradeVolume struct {
	AccountId         Id      `json:"account_id"`
	Symbol            string  `json:"symbol"`
	BaseCurrency      string  `json:"base_currency"`
	NotionalCurrency  string  `json:"notional_currency"`
	DataDate          string  `json:"data_date"`
	TotalVolumeBase   float64 `json:"total_volume_base"`
	MakeBuySellRatio  float64 `json:"maker_buy_sell_ratio"`
	BuyMakerBase      float64 `json:"buy_maker_base"`
	BuyMakerNotional  float64 `json:"buy_maker_notional"`
	BuyMakerCount     float64 `json:"buy_maker_count"`
	SellMakerBase     float64 `json:"sell_maker_base"`
	SellMakerNotional float64 `json:"sell_maker_notional"`
	SellMakerCount    float64 `json:"sell_maker_count"`
	BuyTakerBase      float64 `json:"buy_taker_base"`
	BuyTakerNotional  float64 `json:"buy_taker_notional"`
	BuyTakerCount     float64 `json:"buy_taker_count"`
	SellTakerBase     float64 `json:"sell_taker_base"`
	SellTakerNotional float64 `json:"sell_taker_notional"`
	SellTakerCount    float64 `json:"sell_taker_count"`
}

type CurrentAuction struct {
	ClosedUntil                  int64   `json:"closed_until_ms"`
	LastAuctionEid               Id      `json:"last_auction_eid"`
	LastAuctionPrice             float64 `json:"last_auction_price,string"`
	LastAuctionQuantity          float64 `json:"last_auction_quantity,string"`
	LastHighestBidPrice          float64 `json:"last_highest_bid_price,string"`
	LastLowestAskPrice           float64 `json:"last_lowest_ask_price,string"`
	MostRecentIndicativePrice    float64 `json:"most_recent_indicative_price,string"`
	MostRecentIndicativeQuantity float64 `json:"most_recent_indicative_quantity,string"`
	MostRecentHighestBidPrice    float64 `json:"most_recent_highest_bid_price,string"`
	MostRecentLowestAskPrice     float64 `json:"most_recent_lowest_ask_price,string"`
	NextUpdate                   int64   `json:"next_update_ms"`
	NextAuction                  int64   `json:"next_auction_ms"`
}

type Auction struct {
	Timestamp       int64   `json:"timestampms"`
	AuctionId       Id      `json:"auction_id"`
	Eid             Id      `json:"eid"`
	EventType       string  `json:"event_type"`
	AuctionResult   string  `json:"auction_result"`
	AuctionPrice    float64 `json:"auction_price,string"`
	AuctionQuantity float64 `json:"auction_quantity,string"`
	HighestBidPrice float64 `json:"highest_bid_price,string"`
	LowestAskPrice  float64 `json:"lowest_ask_price,string"`
}

type CancelResult struct {
	Response
	Details CancelResultDetails
}

type CancelResultDetails struct {
	CancelledOrders []Id
	CancelRejects   []Id
}

type FundBalance struct {
	Type                   string  `json:"type"`
	Currency               string  `json:"currency"`
	Amount                 float64 `json:"amount,string"`
	Available              float64 `json:"available,string"`
	AvailableForWithdrawal float64 `json:"availableForWithdrawal,string"`
}

type WithdrawFundsResult struct {
	Destination string  `json:"destination"`
	TxHash      string  `json:"txHash"`
	Amount      float64 `json:"amount,string"`
}
