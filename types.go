package geminix

type Symbol string

type Currency string

type Network string

type Order struct {
	OrderId           string   `json:"order_id"`
	ClientOrderId     string   `json:"client_order_id"`
	Symbol            string   `json:"symbol"`
	Exchange          string   `json:"exchange"`
	Price             string   `json:"price"`
	AvgExecutionPrice string   `json:"avg_execution_price"`
	Side              string   `json:"side"`
	Type              string   `json:"type"`
	Options           []string `json:"options"`
	Timestamp         string   `json:"timestamp"`
	Timestampms       uint64   `json:"timestampms"`
	IsLive            bool     `json:"is_live"`
	IsCancelled       bool     `json:"is_cancelled"`
	Reason            string   `json:"reason"`
	WasForced         bool     `json:"was_forced"`
	ExecutedAmount    string   `json:"executed_amount"`
	RemainingAmount   string   `json:"remaining_amount"`
	OriginalAmount    string   `json:"original_amount"`
	IsHidden          bool     `json:"is_hidden"`
	Trades            []Trade  `json:"trades"`
}

type Trade struct {
	Price         string `json:"price"`
	Amount        string `json:"amount"`
	Timestamp     uint64 `json:"timestamp"`
	Timestampms   uint64 `json:"timestampms"`
	Type          string `json:"type"`
	Aggressor     bool   `json:"aggressor"`
	FeeCurrency   string `json:"fee_currency"`
	FeeAmount     string `json:"fee_amount"`
	Tid           uint   `json:"tid"`
	OrderId       string `json:"order_id"`
	ClientOrderId string `json:"client_order_id"`
	Exchange      string `json:"exchange"`
	IsAuctionFill bool   `json:"is_auction_fill"`
	Break         string `json:"break"`
}

type Balance struct {
	Currency                       Currency `json:"currency"`
	Amount                         string   `json:"amount"`
	AmountNotional                 string   `json:"amountNotional"`
	Available                      string   `json:"available"`
	AvailableNotional              string   `json:"availableNotional"`
	AvailableForWithdrawal         string   `json:"availableForWithdrawal"`
	AvailableForWithdrawalNotional string   `json:"availableForWithdrawalNotional"`
	Type                           string   `json:"type"`
}

type Transfer struct {
	Type        string   `json:"type"`
	Status      string   `json:"status"`
	Timestampms uint64   `json:"timestampms"`
	EID         uint     `json:"eid"`
	AdvanceEid  uint     `json:"advanceEid"`
	Currency    Currency `json:"currency"`
	Amount      string   `json:"amount"`
	Method      string   `json:"method"`
	TxHash      string   `json:"txHash"`
	OutputIdx   uint     `json:"outputIdx"`
	Destination string   `json:"destination"`
	Purpose     string   `json:"purpose"`
}

type CryptoWithdrawal struct {
	Address      string `json:"address"`
	Amount       string `json:"amount"`
	TxHash       string `json:"txHash"`
	WithdrawalId string `json:"withdrawalID"`
	Result       string `json:"result"`
	Reason       string `json:"reason"`
	Message      string `json:"message"`
}

type DepositAddress struct {
	Address   string `json:"address"`
	Timestamp uint64 `json:"timestamp"`
	Label     string `json:"label"`
}

type InternalTransfer struct {
	FromAccount  string   `json:"fromAccount"`
	ToAccount    string   `json:"toAccount"`
	Amount       string   `json:"amount"`
	Fee          string   `json:"fee"`
	Currency     Currency `json:"currency"`
	WithdrawalId string   `json:"withdrawalId"`
	UUID         string   `json:"uuid"`
	Message      string   `json:"message"`
	TxHash       string   `json:"txHash"`
}

type Account struct {
	Name           string `json:"name"`
	AccountName    string `json:"accountName"`
	Account        string `json:"account"`
	ShortName      string `json:"shortName"`
	CounterpartyId string `json:"counterparty_id"`
	Type           string `json:"type"`
	Created        uint64 `json:"created"`
}

type User struct {
	Name        string `json:"name"`
	LastSignIn  string `json:"lastSignIn"`
	Status      string `json:"status"`
	CountryCode string `json:"countryCode"`
	IsVerified  string `json:"isVerified"`
}

type AccountDetail struct {
	Account Account `json:"account"`
	Users   []User  `json:"users"`
}

/*

type Id string

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
}*/
