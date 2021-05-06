package geminix

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

const (
	// fiat
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	SGD Currency = "SGD"

	// crypto
	BTC     Currency = "BTC"
	ETH     Currency = "ETH"
	ZEC     Currency = "ZEC"
	BCH     Currency = "BCH"
	LTC     Currency = "LTC"
	BAT     Currency = "BAT"
	DAI     Currency = "DAI"
	LINK    Currency = "LINK"
	OXT     Currency = "OXT"
	COMP    Currency = "COMP"
	PAXG    Currency = "PAXG"
	MKR     Currency = "MKR"
	ZRX     Currency = "ZRX"
	KNC     Currency = "KNC"
	MANA    Currency = "MANA"
	STORJ   Currency = "STORJ"
	SNX     Currency = "SNX"
	CRV     Currency = "CRV"
	BAL     Currency = "BAL"
	UNI     Currency = "UNI"
	REN     Currency = "REN"
	UMA     Currency = "UMA"
	YFI     Currency = "YFI"
	AAVE    Currency = "AAVE"
	FIL     Currency = "FIL"
	SKL     Currency = "SKL"
	GRT     Currency = "GRT"
	BNT     Currency = "BNT"
	ONEINCH Currency = "1INCH"
	ENJ     Currency = "ENJ"
	LRC     Currency = "LRC"
	SAND    Currency = "SAND"
	CUBE    Currency = "CUBE"
	LPT     Currency = "LPT"
	BOND    Currency = "BOND"
	MATIC   Currency = "MATIC"
	INJ     Currency = "INJ"
	SUSHI   Currency = "SUSHI"
	DOGE    Currency = "DOGE"
)

const (
	Bitcoin     Network = "bitcoin"
	Ethereum    Network = "ethereum"
	BitcoinCash Network = "bitcoincash"
	Litecoin    Network = "litecoin"
	ZCash       Network = "zcash"
	Filecoin    Network = "filecoin"
)
