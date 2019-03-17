package client

import (
	"fmt"
	"net/http"

	"github.com/isfonzar/pagarme-go/internal/payables"
	"github.com/isfonzar/pagarme-go/internal/transactions"
)

func (c *PagarmeClient) CreateTransaction(params transactions.CreateTransactionParams) (*transactions.Transaction, *ErrorResponse, error) {
	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", baseURL, "transactions"), params)
	response := &transactions.Transaction{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

func (c *PagarmeClient) CaptureTransaction(transactionID int, params transactions.CaptureTransactionParams) (*transactions.Transaction, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s/%d/%s", baseURL, "transactions", transactionID, "capture")
	req, err := c.NewRequest(http.MethodPost, route, params)
	response := &transactions.Transaction{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

func (c *PagarmeClient) RefundTransaction(transactionID int, params transactions.RefundTransactionParams) (*transactions.Transaction, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s/%d/%s", baseURL, "transactions", transactionID, "refund")
	req, err := c.NewRequest(http.MethodPost, route, params)
	response := &transactions.Transaction{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

// GetTransactions retorna um slice contendo objetos de transações, ordenadas a partir da transação realizada mais recentemente.
func (c *PagarmeClient) GetTransactions(params transactions.GetTransactionsParams) (*[]transactions.Transaction, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s", baseURL, "transactions")
	req, err := c.NewRequest(http.MethodGet, route, params)
	response := &[]transactions.Transaction{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

// GetTransaction retorna os dados de uma transação em específico, com as informações em um único objeto.
func (c *PagarmeClient) GetTransaction(transactionID int) (*transactions.Transaction, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s/%d", baseURL, "transactions", transactionID)
	req, err := c.NewRequest(http.MethodGet, route, nil)
	response := &transactions.Transaction{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

// GetPayables retorna um slice com objetos recebíveis (payables).
// Os recebíveis são os dados de pagamento referentes a uma transação.
func (c *PagarmeClient) GetPayables(transactionId int) (*[]payables.Payable, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s/%d/%s", baseURL, "transactions", transactionId, "payables")
	req, err := c.NewRequest(http.MethodGet, route, nil)
	response := &[]payables.Payable{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}

// GetPayables retorna um objeto recebível (payable) informando os dados de um pagamento referente a uma transação em específico.
func (c *PagarmeClient) GetPayable(transactionId, payableID int) (*payables.Payable, *ErrorResponse, error) {
	route := fmt.Sprintf("%s/%s/%d/%s/%d", baseURL, "transactions", transactionId, "payables", transactionId)
	req, err := c.NewRequest(http.MethodGet, route, nil)
	response := &payables.Payable{}
	if err != nil {
		return response, nil, err
	}
	errorResp, err := c.Send(req, response)

	return response, errorResp, err
}
