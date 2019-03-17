package transactions

// Os dados sobre endereço devem ser fornecidos dentro do objeto address.
// O campo address é obrigatório para todas as transações de cartão de crédito de companhias com antifraude habilitado.
type Address struct {
	// Obrigatório. Rua
	Street string `json:"street"`
	// Obrigatório. Número
	StreetNumber string `json:"street_number"`
	// Obrigatório. CEP. Para endereço brasileiro, deve conter uma numeração de 8 dígitos
	Zipcode string `json:"zipcode"`
	// Obrigatório. País. Duas letras minúsculas. Deve seguir o padão ISO 3166-1 alpha-2
	Country string `json:"country"`
	// Obrigatório. Estado
	State string `json:"state"`
	// Obrigatório. Cidade
	City string `json:"city"`
	// Bairro
	Neighborhood string `json:"neighborhood"`
	// Complemento. Não pode ser uma string vazia nem null
	Complementary string `json:"complementary"`
}
