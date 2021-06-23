package geminix

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	url    string
	key    string
	secret string
}

func NewClient(key string, secret string, sandbox bool) *Client {
	var url string
	if sandbox {
		url = SandboxBaseUrl
	} else {
		url = BaseUrl
	}

	return &Client{url: url, key: key, secret: secret}
}

func (c *Client) BuildHeader(req *map[string]interface{}) (http.Header, error) {

	reqStr, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	payload := base64.StdEncoding.EncodeToString([]byte(reqStr))

	mac := hmac.New(sha512.New384, []byte(c.secret))
	mac.Write([]byte(payload))

	signature := hex.EncodeToString(mac.Sum(nil))

	header := http.Header{}
	header.Set("X-GEMINI-APIKEY", c.key)
	header.Set("X-GEMINI-PAYLOAD", payload)
	header.Set("X-GEMINI-SIGNATURE", signature)

	return header, nil
}

type Response struct {
	Result string
	ApiError
}

type ApiError struct {
	Reason  string
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("[%v] %v", e.Reason, e.Message)
}

func (c *Client) Request(verb string, uri string, params map[string]interface{}) ([]byte, error) {
	url := c.url + uri

	req, err := http.NewRequest(verb, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	if params != nil {
		if verb == "GET" {
			q := req.URL.Query()
			for key, val := range params {
				q.Add(key, val.(string))
			}
			req.URL.RawQuery = q.Encode()
		} else {
			req.Header, err = c.BuildHeader(&params)
			if err != nil {
				return nil, err
			}
		}
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res Response
	json.Unmarshal(body, &res)
	if res.Result == "error" {
		return nil, &res.ApiError
	}

	return body, nil
}

func (c *Client) PublicRequest(uri string) ([]byte, error) {
	body, err := c.Request("GET", uri, nil)

	return body, err
}

func Nonce() int64 {
	return time.Now().UnixNano()
}

func (c *Client) PrivateRequest(uri string, params map[string]interface{}) ([]byte, error) {
	if params == nil {
		params = map[string]interface{}{
			"request": uri,
			"nonce":   Nonce(),
		}
	} else {
		params["request"] = uri
		params["nonce"] = Nonce()

		/*for key, value := range params {
			if value == nil {
				delete(params, key)
			}
		}*/
	}

	body, err := c.Request("POST", uri, params)
	return body, err
}
