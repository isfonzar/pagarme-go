package bank_accounts

// BankAccount: Objeto de resposta do bank_account
type BankAccount struct {
	// Nome do tipo do objeto criado/modificado. Valor retornado: bank_account
	Object string `json:"object"`
	// Identificador da conta bancária
	ID int `json:"id"`
	// Código do banco
	BankCode string `json:"bank_code"`
	// Agência da conta
	Agencia string `json:"agencia"`
	// Digito verificador da agência
	AgenciaDV string `json:"agencia_dv"`
	// Número da conta bancária
	Conta string `json:"conta"`
	// Dígito verificador da conta
	ContaDV string `json:"conta_dv"`
	// Tipo da conta bancária. Valores possíveis: conta_corrente, conta_poupanca, conta_corrente_conjunta, conta_poupanca_conjunta
	Type string `json:"type"`
	// Tipo do document_number. Valores possíveis: cpf, cnpj
	DocumentType string `json:"document_type"`
	// Número do documento do titular da conta
	DocumentNumber string `json:"document_number"`
	// Nome completo/Razão social do detentor da conta, Até 30 caracteres
	LegalName string `json:"legal_name"`
	// Data de criação da conta bancária (ISODate)
	DateCreated string `json:"date_created"`
}
