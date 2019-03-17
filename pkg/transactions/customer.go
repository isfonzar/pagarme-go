package transactions

type Customer struct {
	// Identificador do cliente em sua plataforma
	ExternalID string `json:"external_id"`
	// Nome do cliente
	Name string `json:"name"`
	// E-mail do cliente
	Email string `json:"email"`
	// Nacionalidade do cliente (br, en ...)
	Country string `json:"country"`
	// Tipo de pessoa, individual ou corporation
	Type string `json:"type"`
	// Dados dos documentos dos clientes
	Documents []Document `json:"documents"`
	// Números de telefone dos clientes
	PhoneNumbers []string `json:"phone_numbers"`
}
