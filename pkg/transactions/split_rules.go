package transactions

type SplitRules struct {
	Liable              bool   `json:"liable"`
	ChargeProcessingFee bool   `json:"charge_processing_fee"`
	Percentage          string `json:"percentage"`
	Amount              string `json:"amount"`
	ChargeRemainderFee  bool   `json:"charge_remainder_fee"`
	RecipientID         string `json:"recipient_id"`
}
