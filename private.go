package geminix

import (
	"encoding/json"
	"fmt"
)

func (c *Client) Balances(account string) ([]Balance, error) {
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

func (c *Client) NotationalBalances(currency Currency, account string) ([]Balance, error) {
	uri := fmt.Sprintf(NotationalBalancesUri, currency)

	params := map[string]interface{}{
		"account": account,
	}

	var notationalBalances []Balance

	response, err := c.PrivateRequest(uri, params)
	if err != nil {
		return notationalBalances, err
	}

	err = json.Unmarshal(response, &notationalBalances)

	return notationalBalances, err
}

func (c *Client) DepositAddresses(network Network, account string) ([]DepositAddress, error) {
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

func (c *Client) AccountDetail(account string) (AccountDetail, error) {
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
