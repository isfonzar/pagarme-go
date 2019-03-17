package client

import (
	"net/http"
	"os"
	"testing"

	"github.com/isfonzar/pagarme-go/internal/transactions"
)

func TestPagarmeClient_CreateTransaction(t *testing.T) {
	address := transactions.Address{
		Street:        "Avenida Brigadeiro Faria Lima",
		StreetNumber:  "1811",
		Neighborhood:  "Jardim Paulistano",
		Zipcode:       "01451001",
		Country:       "br",
		State:         "sp",
		City:          "Sao Paulo",
		Complementary: "Quinto andar",
	}

	params := transactions.CreateTransactionParams{
		Amount:             100,
		CardHolderName:     "Morpheus Fishburne",
		CardExpirationDate: "1220",
		CardNumber:         "4242424242424242",
		CardCVV:            "123",
		PostbackURL:        "http://requestb.in/pkt7pgpk",
		Installments:       "1",
		Billing: transactions.Billing{
			Name:    "Morpheus Fishburne",
			Address: address,
		},
		Shipping: transactions.Shipping{
			Name:         "Neo Reeves",
			DeliveryDate: "2000-12-21",
			Address:      address,
		},
		Items: []transactions.Item{
			{
				Id:        "r123",
				Title:     "Red Pill",
				UnitPrice: 10000,
				Quantity:  1,
				Category:  "Pills",
				Venue:     "Matrix",
				Date:      "2000-12-21",
			},
			{
				Id:        "b123",
				Title:     "Blue Pill",
				UnitPrice: 10000,
				Quantity:  1,
				Category:  "Pills",
				Venue:     "Matrix",
				Date:      "2000-12-21",
			},
		},
		Customer: transactions.Customer{
			ExternalID: "#3311",
			Type:       "individual",
			Name:       "Morpheus Fishburne",
			Country:    "br",
			Email:      "mopheus@nabucodonozor.com",
			Documents: []transactions.Document{
				{
					Number: "30621143049",
					Type:   "cpf",
				},
			},
			PhoneNumbers: []string{
				"+5511999998888",
				"+5511888889999",
			},
		},
	}

	cli := NewClient(http.Client{}, os.Getenv("PAGARME_API_KEY"), nil)

	_, _, err := cli.CreateTransaction(params)

	if err != nil {
		t.Errorf("Error when sending create transaction request got: %s", err.Error())
	}
}

func TestPagarmeClient_CaptureTransaction(t *testing.T) {
	params := transactions.CaptureTransactionParams{
		Amount: 100,
	}

	cli := NewClient(http.Client{}, os.Getenv("PAGARME_API_KEY"), nil)

	_, _, err := cli.CaptureTransaction(12345, params)

	if err != nil {
		t.Errorf("Error when sending capture transaction request got: %s", err.Error())
	}
}

func TestPagarmeClient_RefundTransaction(t *testing.T) {
	params := transactions.NewRefundTransactionParams()

	cli := NewClient(http.Client{}, os.Getenv("PAGARME_API_KEY"), nil)

	_, _, err := cli.RefundTransaction(12345, params)

	if err != nil {
		t.Errorf("Error when sending refund transaction request got: %s", err.Error())
	}
}
