package transactions

// Os dados sobre os produtos vendidos devem ser fornecidos dentro do objeto items.
// O campo items é obrigatório para todas as transações de cartão de crédito de companhias com antifraude habilitado
type Item struct {
	// Obrigatório. SKU (unidade de manutenção de estoque) ou número de identificação na loja
	Id string `json:"id"`
	// Obrigatório. Nome do item vendido.
	Title string `json:"title"`
	// Obrigatório. Preço por unidade. Por exemplo, se o preço de cada item é vinte reais e seis centavos (R$20,06), o valor deve ser fornecido como ‘2006’
	UnitPrice int `json:"unit_price"`
	// Obrigatório. Número de unidades vendidas do produto
	Quantity int `json:"quantity"`
	// Obrigatório. Caracteriza o produto como bem físico ou não.
	Tangible bool `json:"tangible"`
	// Category
	Category string `json:"category"`
	// Local (se evento).
	Venue string `json:"venue"`
	// Data (se evento). Estimativa fornecida no formato AAAA-MM-DD
	Date string `json:"date"`
}
