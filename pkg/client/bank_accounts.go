package client

import (
	"fmt"
	"github.com/isfonzar/pagarme-go/internal/bank_accounts"
)

// CreateBankAccount cria uma conta bancária para futuros pagamentos.
func (c *PagarmeClient) CreateBankAccount(params bank_accounts.CreateBankAccountParams) (*bank_accounts.BankAccount, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s", baseURL, "bank_accounts")
	req, err := c.NewRequest(http.MethodPost, route, params)
	response := &bank_accounts.BankAccount{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

// GetBankAccount: Através dessa rota você consegue retornar os dados de uma conta bancária específica.
func (c *PagarmeClient) GetBankAccount(bankAccountID int) (*bank_accounts.BankAccount, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s/%d", baseURL, "bank_accounts", bankAccountID)
	req, err := c.NewRequest(http.MethodPost, route, nil)
	response := &bank_accounts.BankAccount{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

// GetBankAccount: Através dessa rota você consegue retornar os dados de uma conta bancária específica.
func (c *PagarmeClient) GetBankAccounts(params bank_accounts.GetBankAccountsParams) (*[]bank_accounts.BankAccount, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s", baseURL, "bank_accounts")
	req, err := c.NewRequest(http.MethodGet, route, params)
	response := &[]bank_accounts.BankAccount{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}
