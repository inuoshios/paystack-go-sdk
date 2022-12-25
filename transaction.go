package paystack

import (
	"encoding/json"
	"fmt"
)

var (
	jsonResponse Response
)

type InitTrnx struct {
	Amount            string   `json:"amount,omitempty"`
	Email             string   `json:"email,omitempty"`
	Currency          string   `json:"currency,omitempty"`
	Reference         string   `json:"reference,omitempty"`
	CallbackURL       string   `json:"callback_url,omitempty"`
	Plan              string   `json:"plan,omitempty"`
	InvoiceLimit      uint     `json:"invoice_limit,omitempty"`
	Metadata          string   `json:"metadata,omitempty"`
	Channels          []string `json:"channels,omitempty"`
	SplitCode         string   `json:"split_code,omitempty"`
	Subaccount        string   `json:"subaccount,omitempty"`
	TransactionCharge uint     `json:"transaction_charge,omitempty"`
	Bearer            string   `json:"bearer,omitempty"`
}

func (c *Config) InitializeTransaction(trnx *InitTrnx) (Response, error) {
	path := "/transaction/initialize"

	response, err := c.makeRequest("POST", path, trnx)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, nil
}

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
