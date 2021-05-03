package geminix

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Past Trades
func (c *Client) PastTrades(symbol string, limitTrades int, timestamp int64) ([]Trade, error) {
	uri := PastTradesUri
	url := c.url + uri

	params := map[string]interface{}{
		"request":      uri,
		"nonce":        nonce(),
		"symbol":       symbol,
		"limit_trades": limitTrades,
		"timestamp":    timestamp,
	}

	var trades []Trade

	body, err := c.request("POST", url, params)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &trades)

	return trades, nil
}

// Trade Volume
func (c *Client) TradeVolume() ([][]TradeVolume, error) {
	uri := TradeVolumeUri
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
	}

	var volumes [][]TradeVolume

	body, err := c.request("POST", url, params)
	if err != nil {
		return volumes, err
	}

	json.Unmarshal(body, &volumes)

	return volumes, nil
}

// Active Orders
func (c *Client) ActiveOrders() ([]Order, error) {
	uri := ActiveOrdersUri
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
	}

	var orders []Order

	body, err := c.request("POST", url, params)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &orders)

	return orders, nil
}

// Order Status
func (c *Client) OrderStatus(orderId string) (Order, error) {
	uri := OrderStatusUri
	url := c.url + uri

	params := map[string]interface{}{
		"request":  uri,
		"nonce":    nonce(),
		"order_id": orderId,
	}

	var order Order

	body, err := c.request("POST", url, params)
	if err != nil {
		return order, err
	}

	json.Unmarshal(body, &order)

	return order, nil
}

// New Order
func (c *Client) NewOrder(symbol, clientOrderId string, amount, price float64, side string, options []string) (Order, error) {
	uri := NewOrderUri
	url := c.url + uri

	params := map[string]interface{}{
		"request":         uri,
		"nonce":           nonce(),
		"client_order_id": clientOrderId,
		"symbol":          symbol,
		"amount":          strconv.FormatFloat(amount, 'f', -1, 64),
		"price":           strconv.FormatFloat(price, 'f', -1, 64),
		"side":            side,
		"type":            "exchange limit",
	}

	if options != nil {
		params["options"] = options
	}

	var order Order

	body, err := c.request("POST", url, params)
	if err != nil {
		return order, err
	}

	json.Unmarshal(body, &order)

	return order, nil
}

// Cancel Order
func (c *Client) CancelOrder(orderId string) (Order, error) {
	uri := CancelOrderUri
	url := c.url + uri

	params := map[string]interface{}{
		"request":  uri,
		"nonce":    nonce(),
		"order_id": orderId,
	}

	var order Order

	body, err := c.request("POST", url, params)
	if err != nil {
		return order, err
	}

	json.Unmarshal(body, &order)

	return order, nil
}

// Cancel All
func (c *Client) CancelAll() (CancelResult, error) {
	uri := CancelAllUri
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
	}

	var res CancelResult

	body, err := c.request("POST", url, params)
	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Cancel Session
func (c *Client) CancelSession() (Response, error) {
	uri := CancelSessionUri
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
	}

	var res Response

	body, err := c.request("POST", url, params)
	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Heartbeat
func (c *Client) Heartbeat() (Response, error) {
	uri := HeartbeatUri
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
	}

	var res Response

	body, err := c.request("POST", url, params)
	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Balances
func (c *Client) Balances() ([]FundBalance, error) {
	uri := BalancesUri
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
	}

	var balances []FundBalance

	body, err := c.request("POST", url, params)
	if err != nil {
		return balances, err
	}

	json.Unmarshal(body, &balances)

	return balances, nil
}

// New Deposit Address
func (c *Client) NewDepositAddress(currency, label string) (DepositAddress, error) {
	uri := fmt.Sprintf(NewDepositAddressUri, currency)
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
		"label":   label,
	}

	var res DepositAddress

	body, err := c.request("POST", url, params)
	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Withdraw Crypto Funds
func (c *Client) WithdrawFunds(currency, address string, amount float64) (WithdrawFundsResult, error) {
	uri := fmt.Sprintf(WithdrawCryptoUri, currency)
	url := c.url + uri

	params := map[string]interface{}{
		"request": uri,
		"nonce":   nonce(),
		"address": address,
		"amount":  strconv.FormatFloat(amount, 'f', -1, 64),
	}

	var res WithdrawFundsResult

	body, err := c.request("POST", url, params)
	if err != nil {
		return res, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}
