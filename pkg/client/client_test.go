package client

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	testAPIKey := "test_api_key"
	c := NewClient(testAPIKey, nil)

	if c.apiKey != testAPIKey {
		t.Errorf("Client API key is set incorrectly, got: %s, want: %s", c.apiKey, testAPIKey)
	}

	if c.headers != nil {
		t.Errorf("Client's headers are set incorrectly, got: %v, want: %v", c.headers, nil)
	}

	testHeaders := &Headers{
		"X-PagarMe-User-Agent": "Pagar.me Checkout/2.0.0",
		"X-PagarMe-Version":    "2017-08-28",
	}

	c2 := NewClient(testAPIKey, testHeaders)

	if !reflect.DeepEqual(c2.headers, testHeaders) {
		t.Errorf("Client's headers are set incorrectly, got: %v, want: %v", c2.headers, nil)
	}
}
