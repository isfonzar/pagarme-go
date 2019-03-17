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
