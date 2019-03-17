package transactions

// TransactionSplitRules define as regras de split para a transação
type TransactionSplitRules struct {
	Liable              bool   `json:"liable"`
	ChargeProcessingFee bool   `json:"charge_processing_fee"`
	Percentage          string `json:"percentage"`
	Amount              string `json:"amount"`
	ChargeRemainderFee  bool   `json:"charge_remainder_fee"`
	RecipientID         string `json:"recipient_id"`
}

// RefundSplitRules define as regras de split para o estorno.
type RefundSplitRules struct {
	ID                  string `json:"id"`
	Percentage          string `json:"percentage"`
	Amount              int    `json:"amount"`
	RecipientID         string `json:"recipient_id"`
	ChargeProcessingFee bool   `json:"charge_processing_fee"`
}
