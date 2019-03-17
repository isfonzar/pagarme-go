package transactions

import "github.com/isfonzar/pagarme-go/internal/bank_accounts"

// CreateTransactionParams
type CreateTransactionParams struct {
	// Obrigatório. Valor a ser cobrado. Deve ser passado em centavos. Ex: R$ 10.00 = 1000. Deve ser no mínimo 1 real (100)
	Amount int `json:"amount,omitempty"`
	// Informações do cartão do cliente criptografadas em sua aplicação.
	CardHash string `json:"card_hash,omitempty"`
	// Ao realizar uma transação, retornamos o card_id do cartão,
	// para que nas próximas transações ele possa ser utilizado como forma de identificar os dados de pagamento.
	CardID string `json:"card_id,omitempty"`
	// Nome do portador do cartão.
	CardHolderName string `json:"card_holder_name,omitempty"`
	// Data de validade do cartão no formato MMAA.
	CardExpirationDate string `json:"card_expiration_date,omitempty"`
	// Número do cartão.
	CardNumber string `json:"card_number,omitempty"`
	// Código verificador do cartão.
	CardCVV string `json:"card_cvv,omitempty"`
	// Método de pagamento da transação. Aceita dois tipos: credit_card e boleto
	PaymentMethod string `json:"payment_method,omitempty"`
	// Endpoint do seu sistema que receberá informações a cada atualização da transação.
	// Caso você defina este parâmetro, o processamento da transação se torna assíncrono.
	PostbackURL string `json:"postback_url,omitempty"`
	// Utilize false caso queira manter o processamento síncrono de uma transação.
	Async bool `json:"async,omitempty"`
	// Número de parcelas da transação, sendo mínimo: 1 e Máximo: 12. OBS: Se o pagamento for boleto, o padrão é 1
	Installments string `json:"installments,omitempty"`
	// Prazo limite para pagamento do boleto. Deve ser passado no formato yyyy-MM-dd.Default: data atual + 7 dias
	BoletoExpirationDate string `json:"boleto_expiration_date,omitempty"`
	// Descrição que aparecerá na fatura depois do nome de sua empresa.
	// Máximo de 13 caracteres, sendo alfanuméricos e espaços.
	SoftDescriptor string `json:"soft_descriptor,omitempty"`
	// Após a autorização de uma transação, você pode escolher se irá capturar ou adiar a captura do valor.
	// Caso opte por postergar a captura, atribua o valor false
	Capture string `json:"capture,omitempty"`
	// Campo instruções do boleto. Máximo de 255 caracteres
	BoletoInstructions string `json:"boleto_instructions,omitempty"`
	// Regras de divisão da transação
	SplitRules []TransactionSplitRules `json:"split_rules,omitempty"`
	// Cliente
	Customer Customer `json:"customer,omitempty"`
	// Obrigatório com o antifraude habilitado. Define os dados de cobrança, como nome e endereço
	Billing Billing `json:"billing,omitempty"`
	// Define os dados de envio da compra, como nome do recebedor, valor do envio e endereço de recebimento.
	// Deve ser preenchido no caso da venda de bem físico
	Shipping Shipping `json:"shipping,omitempty"`
	// Obrigatório com o antifraude habilitado. Define os dados dos itens vendidos, como nome, preço unitário e quantidade
	Items []Item `json:"items,omitempty"`
	// Você pode passar dados adicionais na criação da transação para facilitar uma futura análise de dados.
	// Ex: metadata[ idProduto ]=13933139
	Metadata string `json:"metadata,omitempty"`
	// Valor único que identifica a sessão do usuário acessando o site. Máximo de 100 caracteres
	Session string `json:"session,omitempty"`
	// Data e hora do dispositivo que está efetuando a transação. Deve ser enviado no seguinte formato: yyyy-MM-dd'T'HH:mm:ss'Z
	// Este campo é necessário para transações de mundo físico (com método de captura EMV e Magstripe)
	LocalTime string `json:"local_time,omitempty"`
}

// CaptureTransactionParams: utilizado para capturar uma transação, após aprovada, utilizando o token ou o ID da transação.
type CaptureTransactionParams struct {
	// Obrigatório. Valor a ser cobrado. Deve ser passado em centavos. Ex: R$ 10.00 = 1000. Deve ser no mínimo 1 real (100)
	Amount int32 `json:"amount,omitempty"`
	// Regras de divisão da transação
	SplitRules TransactionSplitRules `json:"split_rules"`
	// Metadata de informações adicionais, caso queira incluir, i.e: Id do pedido, nome do produto, etc.
	Metadata string `json:"metadata"`
}

// RefundTransactionParams
type RefundTransactionParams struct {
	// ID da conta bancária. Necessário apenas para estorno de boleto. É possível passar todos os dados da conta também.
	BankAccountID string `json:"bank_account_id,omitempty"`
	// Objeto bank_account que contém os dados da conta bancária para onde o estorno será feito.
	BankAccount bank_accounts.BankAccount `json:"bank_account,omitempty"`
	// Define as regras de split para o estorno.
	SplitRules RefundSplitRules `json:"split_rules,omitempty"`
	// Dígitos que identificam cada banco.
	// Confira a lista dos bancos aqui: http://www.febraban.org.br/associados/utilitarios/Bancos.asp
	BankCode string `json:"bank_code,omitempty"`
	// Número da agência bancária
	Agencia string `json:"agencia,omitempty"`
	// Dígito verificador da agência. Obrigatório caso o banco o utilize
	AgenciaDV string `json:"agencia_dv,omitempty"`
	// Número da conta
	Conta string `json:"conta,omitempty"`
	// CPF ou CNPJ do favorecido
	DocumentNumber string `json:"document_number,omitempty"`
	// Nome/razão social do favorecido, Até 30 caracteres
	LegalName string `json:"legal_name,omitempty"`
	// Define se a operação deve ser feita de maneira assíncrona ou não.
	// Caso true(default), a reposta de sua request será enviada via post para sua postback_url cadastrada na respectiva transação. Caso false, no response será enviado o status final de refunded.
	Async bool `json:"async"`
	// Tipo da conta bancária. Valores possíveis: conta_corrente, conta_poupanca, conta_corrente_conjunta, conta_poupanca_conjunta
	Type string `json:"type,omitempty"`
	// Você pode passar dados adicionais no estorno da transação para facilitar uma futura análise de dados por seus sistemas.
	Metadata string `json:"metadata,omitempty"`
}

// NewRefundTransactionParams retorna um objeto RefundTransactionParams
func NewRefundTransactionParams() RefundTransactionParams {
	params := new(RefundTransactionParams)
	params.Async = true

	return *params
}

// GetTransactionsParams
type GetTransactionsParams struct {
	// Retorna n objetos de transação
	Count int32 `json:"count,omitempty"`
	// Útil para implementação de uma paginação de resultados
	Page int32 `json:"page,omitempty"`
	// Filtro para um dos status: processing, authorized, paid, refunded, waiting_payment, pending_refund, refused
	Status string `json:"status,omitempty"`
	// Filtro para data de criação
	DateCreated string `json:"date_created,omitempty"`
	// Filtro para data de último update
	DateUpdated string `json:"date_updated,omitempty"`
	// Filtro pelo valor da transação
	Amount string `json:"amount,omitempty"`
	// Filtro de parcelas da transação
	Installments string `json:"installments,omitempty"`
	// Filtro de id
	ID string `json:"id,omitempty"`
	// Filtro de tid
	TID string `json:"tid,omitempty"`
	// Filtro de nsu
	NSU string `json:"nsu,omitempty"`
	// Filtro de nome do detentor do cartão
	CardHolderName string `json:"card_holder_name,omitempty"`
	// Filtro dos últimos 4 dígitos do cartão
	CardLastDigits string `json:"card_last_digits,omitempty"`
	// Filtro da bandeira do cartão: amex, aura, discover, elo, hipercard, jcb, visa, mastercard
	CardBrand string `json:"card_brand,omitempty"`
	// Filtro de postback_url
	PostbackURL string `json:"postback_url,omitempty"`
	// Filtro de método de pagamento: credit_card, boleto
	PaymentMethod string `json:"payment_method,omitempty"`
	// Filtro de método de captura
	CaptureMethod string `json:"capture_method,omitempty"`
	// Filtro de boleto_url
	BoletoURL string `json:"boleto_url,omitempty"`
	// Filtro de score de antifraud
	AntifraudScore string `json:"antifraud_score,omitempty"`
	// Filtro de código de barras de boletos
	BoletoBarcode string `json:"boleto_barcode,omitempty"`
	// Filtro de subscription
	SubscriptionID string `json:"subscription_id,omitempty"`
	// Filtro de todos os atributos do cliente
	Customer Customer `json:"customer,omitempty"`
	// Filtro de todos os atributos de endereço
	Address Address `json:"address,omitempty"`
	// Filtro de todos os atributos do telefone
	Phone string `json:"phone,omitempty"`
	// Filtro de todos os parâmetros de metadata. Depende do metadata.
	Metadata string `json:"metadata,omitempty"`
}
