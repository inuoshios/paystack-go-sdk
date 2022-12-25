package paystack

import (
	"encoding/json"
	"fmt"
)

var (
	jsonResponse Response
)

type TransactionBody struct {
	Amount            string   `json:"amount,omitempty"`
	Email             string   `json:"email,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Reference         string   `json:"reference,omitempty"`
	CallbackURL       string   `json:"callback_url,omitempty"`
	Plan              string   `json:"plan,omitempty"`
	InvoiceLimit      uint64   `json:"invoice_limit,omitempty"`
	Metadata          string   `json:"metadata,omitempty"`
	Channels          []string `json:"channels,omitempty"`
	SplitCode         string   `json:"split_code,omitempty"`
	Subaccount        string   `json:"subaccount,omitempty"`
	TransactionCharge uint64   `json:"transaction_charge,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
}

type ChargeAuthorizationBody struct {
	Amount            string   `json:"amount"`
	Email             string   `json:"email"`
	AuthorizationCode string   `json:"authorization_code"`
	Reference         string   `json:"reference,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Metadata          string   `json:"metadata,omitempty"`
	Channels          []string `json:"channels,omitempty"`
	Subaccount        string   `json:"subaccount,omitempty"`
	TransactionCharge uint64   `json:"transaction_charge,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
	Queue             bool     `json:"queue,omitempty"`
}

// InitializeTransaction initiate a new transaction
//
// Docs: https://paystack.com/docs/api/#transaction-initialize
//
//	client, _ := paystack.NewClient(apiKey)
//	transaction, err := client.InitializeTransaction(transactionBody struct{})
func (c *Config) InitializeTransaction(trnx *TransactionBody) (Response, error) {
	path := "/transaction/initialize"

	response, err := c.makeRequest("POST", path, trnx)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, nil
}

// VerifyTransaction confirms the status of a transaction
//
// Docs: https://paystack.com/docs/api/#transaction-verify
//
//	client, _ := paystack.NewClient(apiKey)
//	transaction, err := client.VerifyTransaction(reference string)
func (c *Config) VerifyTransaction(trnxReference string) (Response, error) {
	path := fmt.Sprintf("/transaction/verify/%s", trnxReference)
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ListTransactions returns the transactions carried out on your integration
//
// Docs: https://paystack.com/docs/api/#transaction-list
//
//	client, _ := paystack.NewClient(apiKey)
//	transaction, err := client.ListTransaction()
func (c *Config) ListTransactions() (Response, error) {
	path := "/transaction"
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// FetchTransaction gets details of a transactionn carried out on your integration
//
// Docs: https://paystack.com/docs/api/#transaction-fetch
//
//	client, _ := paystack.NewClient(apiKey)
//	transaction, err := client.FetchTransaction(transactionID uint64)
func (c *Config) FetchTransaction(transactionID uint64) (Response, error) {
	path := fmt.Sprintf("/transaction/%d", transactionID)
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ChargeAuthorization - All authorizations marked as reusable can be charged with this endpoint whenever you need to receive payments.
//
// Docs: https://paystack.com/docs/api/#transaction-charge-authorization
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ChargeAuthorization(body *ChargeAuthorizationBody struct{})
func (c *Config) ChargeAuthorization(body *ChargeAuthorizationBody) (Response, error) {
	path := "/transaction/charge_authorization"
	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}
