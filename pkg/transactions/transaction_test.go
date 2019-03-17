package transactions

import (
	"testing"

	"github.com/isfonzar/pagarme-go/pkg/client"
)

func TestTransaction_Capture(t *testing.T) {
	transaction := Transaction{
		ID: 6003471,
	}

	c := *client.NewClient("api_key", nil)

	_, err := transaction.Capture(c)
	if err != nil {
		t.Errorf("Error when sending request got: %s", err.Error())
	}
}

func TestTransaction_Refund(t *testing.T) {
	transaction := Transaction{
		ID: 6003471,
	}

	c := *client.NewClient("api_key", nil)

	_, err := transaction.Refund(c)
	if err != nil {
		t.Errorf("Error when sending request got: %s", err.Error())
	}
}

func TestTransaction_GetList(t *testing.T) {
	t.Skip() // Not working with an invalid API Key
	transaction := Transaction{}

	c := *client.NewClient("api_key", nil)

	_, err := transaction.GetList(c)
	if err != nil {
		t.Errorf("Error when sending request got: %s", err.Error())
	}
}
