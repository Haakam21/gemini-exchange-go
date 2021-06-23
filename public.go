package geminix

import (
	"encoding/json"
	"fmt"
)

func (c *Client) Ticker(symbol Symbol) (Ticker, error) {
	uri := fmt.Sprintf(TickerUri, symbol)

	var ticker Ticker

	response, err := c.PublicRequest(uri)
	if err != nil {
		return ticker, err
	}

	err = json.Unmarshal(response, &ticker)

	return ticker, err
}
