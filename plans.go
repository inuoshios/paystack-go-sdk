// The Plans API allows you create and manage installment payment options on your integration

package paystack

import (
	"encoding/json"
	"fmt"
)

type Plan struct {
	// Name of Plan
	Name string `json:"name"`

	// Amount should be in kobo if currency is NGN,
	// pesewas, if currency is GHS, and cents, if currency is ZAR
	Amount uint64 `json:"amount"`

	// Interval in words. Valid intervals are: daily, weekly, monthly,biannually, annually.
	Interval string `json:"interval"`

	// A description for this plan
	Description string `json:"description"`

	// Set to false if you don't want invoices to be sent to your customers
	SendInvoices bool `json:"send_invoices"`

	// Set to false if you don't want text messages to be sent to your customers
	SendSMS bool `json:"send_sms"`

	// Currency in which amount is set. Allowed values are NGN, GHS, ZAR or USD
	Currency string `json:"currency"`

	// Number of invoices to raise during subscription to this plan.
	// Can be overridden by specifying an invoice_limit while subscribing.
	InvoiceLimit uint64 `json:"invoice_limit"`
}

// CreatePlan creates a plan on your integration
//
// Docs: https://paystack.com/docs/api/#plan-create
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.CreatePlan(body structs{})
func (c *Config) CreatePlan(body *Plan) (Response, error) {
	path := "/plan"

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ListPlans list plans available on your integration
//
// Docs: https://paystack.com/docs/api/#plan-list
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ListPlans()
func (c *Config) ListPlans() (Response, error) {
	path := "/plan"

	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// FetchPlan gets details of a plan on your integration
// by specifying either the ID or code
//
// Docs: https://paystack.com/docs/api/#plan-fetch
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.FetchPlan(codeOrID string)
func (c *Config) FetchPlan(codeOrID string) (Response, error) {
	path := fmt.Sprintf("/plan/%s", codeOrID)

	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// UpdatePlan updates a plan details on your integration
//
// Docs: https://paystack.com/docs/api/#plan-update
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.UpdatePlan(codeOrID string, body structs{})
func (c *Config) UpdatePlan(codeOrID string, body *Plan) (Response, error) {
	path := fmt.Sprintf("/plan/%s", codeOrID)

	response, err := c.makeRequest("PUT", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}
