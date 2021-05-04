package geminix

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Id             string `json:"account"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	CounterpartyId string `json:"counterparty_id"`
	CreatedAt      int64  `json:"created"`
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
	Timestamp uint   `json:"timestamp"`
	Label     string `json:"label,omitempty"`
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
