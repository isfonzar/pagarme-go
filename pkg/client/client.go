package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

const baseURL = "https://api.pagar.me/1"

type (
	CustomHeaders map[string]string

	PagarmeClient struct {
		http    http.Client
		apiKey  string
		headers *CustomHeaders
	}

	ErrorResponse struct {
		Errors []Error `json:"errors"`
		Url    string  `json:"url"`
		Method string  `json:"method"`
	}

	Error struct {
		Type          string `json:"type"`
		ParameterName string `json:"parameter_name"`
		Message       string `json:"message"`
	}
)

func NewClient(client http.Client, apiKey string, customHeaders *CustomHeaders) *PagarmeClient {
	return &PagarmeClient{client, apiKey, customHeaders}
}

func (c *PagarmeClient) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		var b []byte
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buf)
}

func (c *PagarmeClient) Send(req *http.Request, v interface{}) (*ErrorResponse, error) {
	// Setting headers
	req.Header.Set("Content-Type", "application/json")
	if c.headers != nil {
		for k, v := range *c.headers {
			req.Header.Set(k, v)
		}
	}

	q := req.URL.Query()
	q.Add("api_key", c.apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 500 {
		return nil, errors.New("internal server error")
	}

	if resp.StatusCode != 200 {
		errResp := &ErrorResponse{}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, errResp)

		return errResp, err
	}

	return nil, json.NewDecoder(resp.Body).Decode(v)
}
