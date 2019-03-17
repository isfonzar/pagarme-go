package transactions

// Os dados de cobrança devem ser fornecidos dentro do objeto billing.
// O campo billing é obrigatório para todas as transações de cartão de crédito de companhias com antifraude habilitado.
type Billing struct {
	// Obrigatório. Nome da entidade de cobrança
	Name string `json:"name"`
	// Obrigatório. Dados de endereço de cobrança.
	Address Address `json:"address"`
}
