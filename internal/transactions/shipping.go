package transactions

// Os dados de envio devem ser fornecidos dentro do objeto shipping.
// O campo shipping deve ser preenchido para transações de cartão de crédito de companhias com antifraude habilitado
// em que o item vendido é um bem físico
type Shipping struct {
	// Obrigatório. Nome da entidade de cobrança
	Name string `json:"name"`
	// Obrigatório. Taxa de envio cobrada do comprador em centavos
	Fee int `json:"fee"`
	// Data de entrega. Estimativa fornecida no formato AAAA-MM-DD
	DeliveryDate string `json:"delivery_date"`
	// Entrega expressa. Se for entrega expressa, deve conter ‘true’ (sim). Caso contrário, deve conter ‘false’ (não)
	Expedited bool `json:"expedited"`
	// Obrigatório. Dados do endereço de envio.
	Address Address `json:"address"`
}
