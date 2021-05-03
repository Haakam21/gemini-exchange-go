package geminix

import "encoding/json"

type Account struct {
	Id             string `json:"account"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	CounterpartyId string `json:"counterparty_id"`
	CreatedAt      int64  `json:"created"`
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
