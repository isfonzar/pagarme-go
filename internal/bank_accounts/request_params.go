package bank_accounts

// CreateBankAccountParams cria uma conta bancária para futuros pagamentos.
type CreateBankAccountParams struct {
	// Código do banco OBS: Deve conter 3 caracteres, apenas números
	BankCode string `json:"bank_code,omitempty"`
	// Agência onde sua conta foi criada OBS: Limite de 5 caracteres, apenas números
	Agencia string `json:"agencia,omitempty"`
	// Dígito verificador da sua agência OBS: Deve conter 1 dígito, apenas alfanuméricos
	AgenciaDV string `json:"agencia_dv,omitempty"`
	// Número da conta bancária OBS: Limite de 13 caracteres, apenas números
	Conta string `json:"conta,omitempty"`
	// Dígito verificador da conta OBS: Limite de 2 caracteres, apenas alfanuméricos
	// Se a conta não possuir dv, basta não adicionar a chave conta_dv no json da request.
	ContaDV string `json:"conta_dv,omitempty"`
	// Tipo de conta bancária, valores possíveis: conta_corrente, conta_poupanca, conta_corrente_conjunta, conta_poupanca_conjunta
	Type string `json:"type,omitempty"`
	// Documento identificador do titular da conta (cpf ou cnpj) Ex: 35146484252
	DocumentNumber string `json:"document_number,omitempty"`
	// Nome completo (se pessoa física) ou razão social (se pessoa jurídica). Até 30 caractéres
	LegalName string `json:"legal_name,omitempty"`
}

// GetBankAccountsParams cria uma conta bancária para futuros pagamentos.
type GetBankAccountsParams struct {
	// Retorna n objetos de conta bancária
	Count string `json:"count,omitempty"`
	// Útil para implementação de uma paginação de resultados
	Page string `json:"page,omitempty"`
	// Filtro pelo ID da conta bancária
	ID string `json:"id,omitempty"`
	// Filtro pelo código do banco das contas bancárias
	BankCode string `json:"bank_code,omitempty"`
	// Filtro pela agencia da conta
	Agencia string `json:"agencia,omitempty"`
	// Filtro pelo dígito verificador da agencia da conta
	AgenciaDV string `json:"agencia_dv,omitempty"`
	// Filtro pelo número da conta
	Conta string `json:"conta,omitempty"`
	// Filtro pelo dígito verificador da conta
	ContaDV string `json:"conta_dv,omitempty"`
	// Filtro pelo CPF/CNPJ atrelado à conta
	DocumentNumber string `json:"document_number,omitempty"`
	// Filtro pelo nome do detentor da conta
	LegalName string `json:"legal_name,omitempty"`
	// Filtro pela data de criação da conta
	DateCreated string `json:"date_created,omitempty"`
}
