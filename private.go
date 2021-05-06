package geminix

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Name           string `json:"name"`
	AccountName    string `json:"accountName"`
	Account        string `json:"account"`
	ShortName      string `json:"shortName"`
	CounterpartyId string `json:"counterparty_id"`
	Type           string `json:"type"`
	Created        uint64 `json:"created"`
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

func (c *Client) GetAccounts() ([]Account, error) {
	var accounts []Account

	response, err := c.PrivateRequest(AccountsUri, nil)
	if err != nil {
		return accounts, err
	}

	err = json.Unmarshal(response, &accounts)

	return accounts, err
}

type DepositAddress struct {
	Address   string `json:"address"`
	Timestamp uint64 `json:"timestamp"`
	Label     string `json:"label"`
}

func (c *Client) GetDepositAddresses(network Network, accountId string) ([]DepositAddress, error) {
	uri := fmt.Sprintf(DepositAddressesUri, network)

	params := map[string]interface{}{
		"account": accountId,
	}

	var depositAddresses []DepositAddress

	response, err := c.PrivateRequest(uri, params)
	if err != nil {
		return depositAddresses, err
	}

	err = json.Unmarshal(response, &depositAddresses)

	return depositAddresses, err
}
