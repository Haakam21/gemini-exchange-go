package geminix

import (
	"encoding/json"
	"fmt"
)

func (c *Client) NewOrder(clientOrderId *uint, symbol Symbol, amount string, minAmount *string, price string, side string, Type string, options *[]string, stopPrice *string, account *string) (Order, error) {
	params := map[string]interface{}{
		"client_order_id": clientOrderId,
		"symbol":          symbol,
		"amount":          amount,
		"min_amount":      minAmount,
		"price":           price,
		"side":            side,
		"type":            Type,
		"options":         options,
		"stop_price":      stopPrice,
		"account":         account,
	}

	var order Order

	response, err := c.PrivateRequest(NewOrderUri, params)
	if err != nil {
		return order, err
	}

	err = json.Unmarshal(response, &order)

	return order, err
}

func (c *Client) CancelOrder(orderId uint, account *string) (Order, error) {
	params := map[string]interface{}{
		"order_id": orderId,
		"account":  account,
	}

	var order Order

	response, err := c.PrivateRequest(CancelOrderUri, params)
	if err != nil {
		return order, err
	}

	err = json.Unmarshal(response, &order)

	return order, err
}

func (c *Client) OrderStatus(orderId uint, clientOrderId *uint, includeTrades *bool, account *string) (Order, error) {
	params := map[string]interface{}{
		"order_id":        orderId,
		"client_order_id": clientOrderId,
		"include_trades":  includeTrades,
		"account":         account,
	}

	var order Order

	response, err := c.PrivateRequest(OrderStatusUri, params)
	if err != nil {
		return order, err
	}

	err = json.Unmarshal(response, &order)

	return order, err
}

func (c *Client) ActiveOrders(account *string) ([]Order, error) {
	params := map[string]interface{}{
		"account": account,
	}

	var orders []Order

	response, err := c.PrivateRequest(ActiveOrdersUri, params)
	if err != nil {
		return orders, err
	}

	err = json.Unmarshal(response, &orders)

	return orders, err
}

func (c *Client) PastTrades(symbol Symbol, limitTrades *uint, timestamp *uint64, account *string) ([]Trade, error) {
	params := map[string]interface{}{
		"symbol":       symbol,
		"limit_trades": limitTrades,
		"timestamp":    timestamp,
		"account":      account,
	}

	var trades []Trade

	response, err := c.PrivateRequest(PastTradesUri, params)
	if err != nil {
		return trades, err
	}

	err = json.Unmarshal(response, &trades)

	return trades, err
}

func (c *Client) Balances(account *string) ([]Balance, error) {
	params := map[string]interface{}{
		"account": account,
	}

	var balances []Balance

	response, err := c.PrivateRequest(BalancesUri, params)
	if err != nil {
		return balances, err
	}

	err = json.Unmarshal(response, &balances)

	return balances, err
}

func (c *Client) NotionalBalances(currency Currency, account *string) ([]Balance, error) {
	uri := fmt.Sprintf(NotionalBalancesUri, currency)

	params := map[string]interface{}{
		"account": account,
	}

	var notionalBalances []Balance

	response, err := c.PrivateRequest(uri, params)
	if err != nil {
		return notionalBalances, err
	}

	err = json.Unmarshal(response, &notionalBalances)

	return notionalBalances, err
}

func (c *Client) Transfers(timestamp *uint64, limitTransfers *uint, account *string, completedAdvances *bool) ([]Transfer, error) {
	params := map[string]interface{}{
		"timestamp":                       timestamp,
		"limit_transfers":                 limitTransfers,
		"account":                         account,
		"show_completed_deposit_advances": completedAdvances,
	}

	var transfers []Transfer

	response, err := c.PrivateRequest(TransfersUri, params)
	if err != nil {
		return transfers, err
	}

	err = json.Unmarshal(response, &transfers)

	return transfers, err
}

func (c *Client) WithdrawCrypto(currency Currency, address string, amount string, account *string) (CryptoWithdrawal, error) {
	uri := fmt.Sprintf(WithdrawCryptoUri, currency)

	params := map[string]interface{}{
		"address": address,
		"amount":  amount,
		"account": account,
	}

	var cryptoWithdrawal CryptoWithdrawal

	response, err := c.PrivateRequest(uri, params)
	if err != nil {
		return cryptoWithdrawal, err
	}

	err = json.Unmarshal(response, &cryptoWithdrawal)

	return cryptoWithdrawal, err
}

func (c *Client) DepositAddresses(network Network, account *string) ([]DepositAddress, error) {
	uri := fmt.Sprintf(DepositAddressesUri, network)

	params := map[string]interface{}{
		"account": account,
	}

	var depositAddresses []DepositAddress

	response, err := c.PrivateRequest(uri, params)
	if err != nil {
		return depositAddresses, err
	}

	err = json.Unmarshal(response, &depositAddresses)

	return depositAddresses, err
}

func (c *Client) InternalTransfer(currency Currency, sourceAccount string, targetAccount string, amount string) (InternalTransfer, error) {
	uri := fmt.Sprintf(InternalTransferUri, currency)

	params := map[string]interface{}{
		"sourceAccount": sourceAccount,
		"targetAccount": targetAccount,
		"amount":        amount,
	}

	var internalTransfer InternalTransfer

	response, err := c.PrivateRequest(uri, params)
	if err != nil {
		return internalTransfer, err
	}

	err = json.Unmarshal(response, &internalTransfer)

	return internalTransfer, err
}

func (c *Client) AccountDetail(account *string) (AccountDetail, error) {
	params := map[string]interface{}{
		"account": account,
	}

	var accountDetail AccountDetail

	response, err := c.PrivateRequest(AccountDetailUri, params)
	if err != nil {
		return accountDetail, err
	}

	err = json.Unmarshal(response, &accountDetail)

	return accountDetail, err
}

func (c *Client) CreateAccount(name string, Type string) (Account, error) {
	params := map[string]interface{}{
		"name": name,
		"type": Type,
	}

	var account Account

	response, err := c.PrivateRequest(CreateAccountUri, params)
	if err != nil {
		return account, err
	}

	err = json.Unmarshal(response, &account)

	return account, err
}

func (c *Client) Accounts() ([]Account, error) {
	var accounts []Account

	response, err := c.PrivateRequest(AccountsUri, nil)
	if err != nil {
		return accounts, err
	}

	err = json.Unmarshal(response, &accounts)

	return accounts, err
}
