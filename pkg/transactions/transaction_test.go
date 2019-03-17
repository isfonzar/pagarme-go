package transactions

import (
	"testing"

	"github.com/isfonzar/pagarme-go/pkg/client"
)

func TestTransaction_Create(t *testing.T) {
	address := Address{
		Street:        "Avenida Brigadeiro Faria Lima",
		StreetNumber:  "1811",
		Neighborhood:  "Jardim Paulistano",
		Zipcode:       "01451001",
		Country:       "br",
		State:         "sp",
		City:          "Sao Paulo",
		Complementary: "Quinto andar",
	}

	trans := Transaction{
		Amount:             100,
		CardHolderName:     "Morpheus Fishburne",
		CardExpirationDate: "1220",
		CardNumber:         "4242424242424242",
		CardCVV:            "123",
		PostbackURL:        "http://requestb.in/pkt7pgpk",
		Installments:       "1",
		Billing: Billing{
			Name:    "Morpheus Fishburne",
			Address: address,
		},
		Shipping: Shipping{
			Name:         "Neo Reeves",
			DeliveryDate: "2000-12-21",
			Address:      address,
		},
		Items: []Item{
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
		Customer: Customer{
			ExternalID: "#3311",
			Type:       "individual",
			Name:       "Morpheus Fishburne",
			Country:    "br",
			Email:      "mopheus@nabucodonozor.com",
			Documents: []Document{
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

	cli := *client.NewClient("api_key", nil)

	_, err := trans.Create(cli)
	if err != nil {
		t.Errorf("Error when sending request got: %s", err.Error())
	}
}
