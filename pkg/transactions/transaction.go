package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/isfonzar/pagarme-go/pkg/client"
)

const endPoint = "transactions"

// Transaction
type Transaction struct {
	// Obrigatório. Valor a ser cobrado. Deve ser passado em centavos. Ex: R$ 10.00 = 1000. Deve ser no mínimo 1 real (100)
	Amount int `json:"amount"`
	// Informações do cartão do cliente criptografadas em sua aplicação.
	CardHash string `json:"card_hash"`
	// Ao realizar uma transação, retornamos o card_id do cartão,
	// para que nas próximas transações ele possa ser utilizado como forma de identificar os dados de pagamento.
	CardID string `json:"card_id"`
	// Nome do portador do cartão.
	CardHolderName string `json:"card_holder_name"`
	// Data de validade do cartão no formato MMAA.
	CardExpirationDate string `json:"card_expiration_date"`
	// Número do cartão.
	CardNumber string `json:"card_number"`
	// Código verificador do cartão.
	CardCVV string `json:"card_cvv"`
	// Método de pagamento da transação. Aceita dois tipos: credit_card e boleto
	PaymentMethod string `json:"payment_method"`
	// Endpoint do seu sistema que receberá informações a cada atualização da transação.
	// Caso você defina este parâmetro, o processamento da transação se torna assíncrono.
	PostbackURL string `json:"postback_url"`
	// Utilize false caso queira manter o processamento síncrono de uma transação.
	Async bool `json:"async"`
	// Número de parcelas da transação, sendo mínimo: 1 e Máximo: 12. OBS: Se o pagamento for boleto, o padrão é 1
	Installments string `json:"installments"`
	// Prazo limite para pagamento do boleto. Deve ser passado no formato yyyy-MM-dd.Default: data atual + 7 dias
	BoletoExpirationDate string `json:"boleto_expiration_date"`
	// Descrição que aparecerá na fatura depois do nome de sua empresa.
	// Máximo de 13 caracteres, sendo alfanuméricos e espaços.
	SoftDescriptor string `json:"soft_descriptor"`
	// Após a autorização de uma transação, você pode escolher se irá capturar ou adiar a captura do valor.
	// Caso opte por postergar a captura, atribua o valor false
	Capture string `json:"capture"`
	// Campo instruções do boleto. Máximo de 255 caracteres
	BoletoInstructions string `json:"boleto_instructions"`
	// Regras de divisão da transação
	SplitRules []struct {
		Liable              bool   `json:"liable"`
		ChargeProcessingFee bool   `json:"charge_processing_fee"`
		Percentage          string `json:"percentage"`
		Amount              string `json:"amount"`
		ChargeRemainderFee  bool   `json:"charge_remainder_fee"`
		RecipientID         string `json:"recipient_id"`
	} `json:"split_rules"`
	// Cliente
	Customer Customer `json:"customer"`
	// Obrigatório com o antifraude habilitado. Define os dados de cobrança, como nome e endereço
	Billing Billing `json:"billing"`
	// Define os dados de envio da compra, como nome do recebedor, valor do envio e endereço de recebimento.
	// Deve ser preenchido no caso da venda de bem físico
	Shipping Shipping `json:"shipping"`
	// Obrigatório com o antifraude habilitado. Define os dados dos itens vendidos, como nome, preço unitário e quantidade
	Items []Item `json:"items"`
	// Você pode passar dados adicionais na criação da transação para facilitar uma futura análise de dados.
	// Ex: metadata[ idProduto ]=13933139
	Metadata string `json:"metadata"`
	// Valor único que identifica a sessão do usuário acessando o site. Máximo de 100 caracteres
	Session string `json:"session"`
	// Data e hora do dispositivo que está efetuando a transação. Deve ser enviado no seguinte formato: yyyy-MM-dd'T'HH:mm:ss'Z
	// Este campo é necessário para transações de mundo físico (com método de captura EMV e Magstripe)
	LocalTime string `json:"local_time"`
}

// Create cria uma transação
func (t *Transaction) Create(client client.Client) (*http.Response, error) {
	transactionReq, err := json.Marshal(&t)
	if err != nil {
		return nil, err
	}

	return client.Request(http.MethodPost, endPoint, transactionReq)
}
