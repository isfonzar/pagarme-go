package transactions

// Os dados sobre documentos devem ser fornecidos dentro do objeto documents.
// O campo documents é obrigatório para todas as transações de cartão de crédito de companhias com antifraude habilitado.
type Document struct {
	// Obrigatório. Tipo de documento. Para compradores brasileiros, deve ser fornecido ao menos um CPF
	// (no caso de pessoa física, i.e. individual) ou CNPJ (no caso de pessoa jurídica, i.e. corporation).
	// Para compradores internacionais, o documento pode ser um passaporte ou um campo personalizado.
	Type string `json:"type"`
	// Obrigatório. Número do documento
	Number string `json:"number"`
}
