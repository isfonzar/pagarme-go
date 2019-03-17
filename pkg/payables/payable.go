package payables

// Objeto contendo os dados de um recebível. O recebível (payable) é gerado automaticamente após uma transação ser paga.
// Para cada parcela de uma transação é gerado um recebível, que também pode ser dividido por recebedor
// (no caso de um split ter sido feito).
type Payable struct {
	// Nome do tipo do objeto criado/modificado.
	Object string `json:"object"`
	// Identificador do recebível
	ID int `json:"id"`
	// Estado atual do recebível. Valores possíveis: waiting_funds, prepaid, paid, suspended.
	Status string `json:"status"`
	// Valor em centavos que foi pago
	Amount int `json:"amount"`
	// Valor em centavos que foi cobrado (taxa)
	Fee int `json:"fee"`
	// Valor em centavos que foi cobrado pela respectiva antecipação(caso criada)
	AnticipationFee int `json:"anticipation_fee"`
	// Número da parcela
	Installment int `json:"installment"`
	// Identificador da transação que gerou o recebível
	TransactionID int `json:"transaction_id"`
	// Identificador da regra de split do recebível
	SplitRuleID string `json:"split_rule_id"`
	// Identificador da antecipação deste recebível, caso tenha sido criada.
	BulkAnticipationID string `json:"bulk_anticipation_id"`
	// Identificador do recebedor a qual este recebível pertence
	RecipientID string `json:"recipient_id"`
	// Dia e hora do pagamento (ISODate)
	PaymentDate string `json:"payment_date"`
	// Dia de hora de pagamento (ISODate) originais do recebível
	OriginalPaymentDate string `json:"original_payment_date"`
	// Tipo do recebível.
	// Valores possíveis: credit, refund, refund_reversal, chargeback, chargeback_refund, block, unblock.
	Type string `json:"type"`
	// Tipo de pagamento da transação. Valores possíveis: credit_card, debit_card, boleto.
	PaymentMethod string `json:"payment_method"`
	// Data(ISODate) em que o banco emissor reconheceu o pagamento da transação a que este recebível pertence.
	AccrualDate string `json:"accrual_date"`
	// Data da criação do objeto (ISODate)
	DateCreated string `json:"date_created"`
}

type PayableList []Payable
