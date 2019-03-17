package transactions

// Transaction é o objeto que você recebe como resposta em cada etapa do processo de efetivação da transação
type Transaction struct {
	// Object Nome do tipo do objeto criado/modificado.
	Object string `json:"object"`
	// Representa o estado da transação. A cada atualização no processamento da transação, esta propriedade é alterada
	// caso você esteja usando uma postback_url, os seus servidores são notificados desses updates.
	// Valores possíveis: processing, authorized, paid, refunded, waiting_payment, pending_refund, refused .
	Status string `json:"status"`
	// Motivo pelo qual a transação foi recusada.
	// Valores possíveis: acquirer, antifraud, internal_error, no_acquirer, acquirer_timeout
	RefuseReason string `json:"refuse_reason"`
	// Agente responsável pela validação ou anulação da transação.
	// Valores possíveis: acquirer, antifraud, internal_error, no_acquirer, acquirer_timeout
	StatusReason string `json:"status_reason"`
	// Adquirente responsável pelo processamento da transação.
	// Valores possíveis: development (em ambiente de testes), pagarme (adquirente Pagar.me), stone, cielo, rede.
	AcquirerName string `json:"acquirer_name"`
	// ID da adquirente responsável pelo processamento da transação.
	AcquirerID string `json:"acquirer_id"`
	// Mensagem de resposta da adquirente referente ao status da transação.
	AcquirerResponseCode string `json:"acquirer_response_code"`
	// Código de autorização retornado pela bandeira.
	AuthorizationCode string `json:"authorization_code"`
	// Texto que irá aparecer na fatura do cliente depois do nome da loja. OBS: Limite de 13 caracteres.
	SoftDescriptor string `json:"soft_descriptor"`
	// Código que identifica a transação na adquirente.
	TID int `json:"tid"`
	// Código que identifica a transação na adquirente.
	NSU int `json:"nsu"`
	// Data de criação da transação no formato ISODate
	DateCreated string `json:"date_created"`
	// Data de atualização da transação no formato ISODate
	DateUpdated string `json:"date_updated"`
	// Valor, em centavos, da transação.
	Amount int `json:"amount"`
	// Valor em centavos autorizado na transação, sempre menor ou igual a amount.
	AuthorizedAmount int `json:"authorized_amount"`
	// Valor em centavos capturado na transação, sempre menor ou igual a authorized_amount.
	PaidAmount int `json:"paid_amount"`
	// Valor em centavos estornado até o momento na transação, sempre menor ou igual a paidamount
	RefundedAmount int `json:"refunded_amount"`
	// Número de parcelas a serem cobradas. OBS: Mínimo 1 e Máximo 12.
	Installments int `json:"installments"`
	// Número identificador da transação
	ID int `json:"id"`
	// Custo da transação para o lojista, envolvendo processamento e antifraude.
	Cost int `json:"cost"`
	// Nome do portador do cartão.
	CardHolderName string `json:"card_holder_name"`
	// Últimos 4 dígitos do cartão.
	CardLastDigits string `json:"card_last_digits"`
	// Primeiros 5 dígitos do cartão
	CardFirstDigits string `json:"card_first_digits"`
	// Bandeira do cartão.
	CardBrand string `json:"card_brand"`
	// Usado em transações EMV, define se a validação do cartão aconteceu online(com banco emissor), ou offline(através do chip).
	CardPinMode string `json:"card_pin_mode"`
	// URL (endpoint) de seu sistema que recebe notificações a cada mudança no status da transação.
	PostbackURL string `json:"postback_url"`
	// Método de pagamento, com os possíveis valores: credit_card e boleto.
	PaymentMethod string `json:"payment_method"`
	// Define qual foi a forma de captura dos dados de pagamento. Valores possíveis: magstripe, emv, ecommerce.
	CaptureMethod string `json:"capture_method"`
	// Define qual foi a nota de antifraude atribuída a transação. Por padrão, transações com score >= 95 são recusadas.
	AntifraudScore string `json:"antifraud_score"`
	// URL do boleto para impressão
	BoletoURL string `json:"boleto_url"`
	// Código de barras do boleto gerado na transação
	BoletoBarcode string `json:"boleto_barcode"`
	// Data de expiração do boleto (em ISODate)
	BoletoExpirationDate string `json:"boleto_expiration_date"`
	// Mostra se a transação foi criada utilizando a API Key ou Encryption Key
	Referer string `json:"referer"`
	// IP de origem que criou a transação
	// podendo ser diretamente de seu cliente, caso a requisição venha diretamente do client-side,
	// ou de seus servidores, caso tudo esteja centralizando em sua aplicação no server-side.
	IP string `json:"ip"`
	// Caso essa transação tenha sido originada na cobrança de uma assinatura, o id desta será o valor dessa propriedade.
	SubscriptionID int `json:"subscription_id"`
	// Dados do cliente. Obrigatório com o antifraude habilitado. O objeto customer é descrito aqui
	Customer Customer `json:"customer"`
	// Dados de cobrança da transação. Obrigatório com o antifraude habilitado. O objeto billing é descrito aqui
	Billing Billing `json:"billing"`
	// Dados de envio do que foi comprado. Deve ser preenchido no caso de venda de bem físico. O objeto shipping é descrito aqui
	Shipping Shipping `json:"shipping"`
	// Dados sobre os produtos comprados. Obrigatório com o antifraude habilitado. O objeto items é descrito aqui
	Items []Item `json:"items"`
	// Dados de endereço, presente em shipping e billing. Obrigatório com o antifraude habilitado. O objeto address é descrito aqui
	Address Address `json:"address"`
	// Informações de documentos do comprador. Obrigatório com o antifraude habilitado. O objeto documents é descrito aqui
	Documents []Document `json:"documents"`
	// Objeto com dados adicionais informados na criação da transação.
	Metadata struct{} `json:"metadata"`
	// Objeto com as regras de split definidas para essa transação.
	SplitRules []TransactionSplitRules `json:"split_rules,omitempty"`
	// Objeto com dados usados na integração com antifraude.
	AntifraudMetadata struct{} `json:"antifraud_metadata"`
	// Valor único que identifica a sessão do usuário acessando o site
	Session string `json:"session"`
	// Define as regras de split para o estorno.
	RefundSplitRules []RefundSplitRules `json:"split_rules,omitempty"`
}

type TransactionList []Transaction
