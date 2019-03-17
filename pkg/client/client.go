package client

import (
	"bytes"
	"net/http"
)

const baseURI = "https://api.pagar.me/1/"

type Client struct {
	apiKey  string
	headers *Headers
}

func NewClient(apiKey string, headers *Headers) *Client {
	return &Client{apiKey, headers}
}

func (c *Client) Request(method string, route string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, baseURI+route, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// Setting headers
	req.Header.Set("Content-Type", "application/json")
	if c.headers != nil {
		for k, v := range *c.headers {
			req.Header.Set(k, v)
		}
	}

	// Authentication
	q := req.URL.Query()
	q.Add("api_key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	httpClient := &http.Client{}

	return httpClient.Do(req)
}
