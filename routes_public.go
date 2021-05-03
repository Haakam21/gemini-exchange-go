package geminix

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Symbols
func (c *Client) Symbols() ([]string, error) {
	uri := SymbolsUri
	url := c.url + uri

	var symbols []string

	body, err := c.request("GET", url, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &symbols)

	return symbols, nil
}

// Ticker
func (c *Client) Ticker(symbol string) (Ticker, error) {
	uri := fmt.Sprintf(TickerUri, symbol)
	url := c.url + uri

	var ticker Ticker

	body, err := c.request("GET", url, nil)
	if err != nil {
		return ticker, err
	}

	json.Unmarshal(body, &ticker)

	return ticker, nil
}

// Order Book
func (c *Client) OrderBook(symbol string, limitBids, limitAsks int) (Book, error) {
	uri := fmt.Sprintf(OrderBookUri, symbol)
	url := c.url + uri

	params := map[string]interface{}{
		"limit_bids": strconv.Itoa(limitBids),
		"limit_asks": strconv.Itoa(limitAsks),
	}

	var book Book

	body, err := c.request("GET", url, params)
	if err != nil {
		return book, err
	}

	json.Unmarshal(body, &book)

	return book, nil
}

// Trades
func (c *Client) Trades(symbol string, since int64, limitTrades int, includeBreaks bool) ([]Trade, error) {
	uri := fmt.Sprintf(TradesUri, symbol)
	url := c.url + uri

	params := map[string]interface{}{
		"since":          strconv.Itoa(int(since)),
		"limit_trades":   strconv.Itoa(limitTrades),
		"include_breaks": strconv.FormatBool(includeBreaks),
	}

	var res []Trade

	body, err := c.request("GET", url, params)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Current Auction
func (c *Client) CurrentAuction(symbol string) (CurrentAuction, error) {
	uri := fmt.Sprintf(TradesUri, symbol)
	url := c.url + uri

	var auction CurrentAuction

	body, err := c.request("GET", url, nil)
	if err != nil {
		return auction, err
	}

	json.Unmarshal(body, &auction)

	return auction, nil
}

// Auction History
func (c *Client) AuctionHistory(symbol string, since int64, limit int, includeIndicative bool) ([]Auction, error) {
	uri := fmt.Sprintf(AuctionHistoryUri, symbol)
	url := c.url + uri

	params := map[string]interface{}{
		"since":                 strconv.Itoa(int(since)),
		"limit_auction_results": strconv.Itoa(limit),
		"include_indicative":    strconv.FormatBool(includeIndicative),
	}

	var auctions []Auction

	body, err := c.request("GET", url, params)
	if err != nil {
		return auctions, err
	}

	json.Unmarshal(body, &auctions)

	return auctions, nil
}
