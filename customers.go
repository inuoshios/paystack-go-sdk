// The Customers API allows you create and manage customers on your integration.

package paystack

import (
	"encoding/json"
	"fmt"
)

type CreateCustomerBody struct {
	// Email: Customer's email address
	Email string `json:"email"`

	// FirstName: Customer's first name
	FirstName string `json:"first_name,omitempty"`

	// LastName: Customer's last name
	LastName string `json:"last_name,omitempty"`

	// Phone: Customer's phone number
	Phone string `json:"phone,omitempty"`

	// Metadata: A set of key/value pairs that you can attach to the customer.
	//It can be used to store additional information in a structured format.
	Metadata map[string]any
}

type UpdateCustomerBody struct {
	// FirstName: Customer's first name
	FirstName string `json:"first_name"`

	// LastName: Customer's last name
	LastName string `json:"last_name"`

	// Phone: Customer's phone number
	Phone string `json:"phone"`

	// Metadata: A set of key/value pairs that you can attach to the customer.
	//It can be used to store additional information in a structured format.
	Metadata map[string]any
}

type ValidateCustomerBody struct {
	// FirstName: Customer's first name
	FirstName string `json:"first_name"`

	// LastName: Customer's last name
	LastName string `json:"last_name"`

	// Type: Predefined types of identification.
	// Only bank_account is supported at the moment
	Type string `json:"type"`

	// Value: Customer's identification number
	Value string `json:"value"`

	// Country: 2 letter country code of identification issuer
	Country string `json:"country"`

	// BVN: Customer's Bank Verification Number
	BVN string `json:"bvn"`

	// BankCode: You can get the list of Bank Codes by calling
	// the List Banks endpoint. (required if type is bank_account)
	//
	// https://paystack.com/docs/api/#miscellaneous-bank
	BankCode string `json:"bank_code"`

	// AccountNumber: Customer's bank account number. (required if type is bank_account)
	AccountNumber string `json:"account_number"`

	// MiddleName: Customer's middle name - Optional
	MiddleName string `json:"middle_name,omitempty"`
}

type WhiteListOrBlacklistCustomerBody struct {
	// Customer's code, or email address
	Customer string `json:"customer"`

	// One of the possible risk actions [ default, allow, deny ].
	// allow to whitelist. deny to blacklist. Customers start with a default risk action.
	RiskAction string `json:"risk_action"`
}

type DeactivateAuthorizationBody struct {
	// Authorization code to be deactivated
	AuthorizationCode string `json:"authorization_code"`
}

// CreateCustomer create a customer on your integration
//
// **Customer Validation**
//
// The first_name, last_name and phone are optional parameters.
// However, when creating a customer that would be assigned a
// Dedicated Virtual Account and your business catgeory falls
// under Betting, Financial services, and General Service,
// then these parameters become compulsory.
//
// Docs: https://paystack.com/docs/api/#customer-create
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.CreateCustomer(body struct{}})
func (c *Config) CreateCustomer(body *CreateCustomerBody) (Response, error) {
	path := "/customer"

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ListCustomers lists customers available on your integration
//
// Docs: https://paystack.com/docs/api/#customer-list
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ListCustomers()
func (c *Config) ListCustomers() (Response, error) {
	path := "/customer"

	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// FetchCustomer gets details of a customer on your integration
//
// Docs: https://paystack.com/docs/api/#customer-fetch
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.FetchCustomer(emailOrCode string)
func (c *Config) FetchCustomer(emailOrCode string) (Response, error) {
	path := fmt.Sprintf("/customer/%s", emailOrCode)

	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// UpdateCustomer updates a customer's detail on your integration
//
// Docs: https://paystack.com/docs/api/#customer-update
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.UpdateCustomer(code string, body structs{})
func (c *Config) UpdateCustomer(code string, body *UpdateCustomerBody) (Response, error) {
	path := fmt.Sprintf("/customer/%s", code)

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ValidateCustomer validates a customer's identity
//
// Docs: https://paystack.com/docs/api/#customer-validate
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ValidateCustomer(code string, body structs{})
func (c *Config) ValidateCustomer(code string, body *ValidateCustomerBody) (Response, error) {
	path := fmt.Sprintf("/customer/%s/identification", code)

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// WhiteListOrBlacklistCustomer: Whitelist or blacklist a customer on your integration
//
// Docs: https://paystack.com/docs/api/#customer-whitelist-blacklist
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.WhiteListOrBlacklistCustomer(body structs{})
func (c *Config) WhiteListOrBlacklistCustomer(body *WhiteListOrBlacklistCustomerBody) (Response, error) {
	path := "/customer/set_risk_action"

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// DeactivateAuthorization: Deactivate an authorization when the card needs to be forgotten
//
// Docs: https://paystack.com/docs/api/#customer-deactivate-authorization
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.DeactivateAuthorization(body structs{})
func (c *Config) DeactivateAuthorization(body *DeactivateAuthorizationBody) (Response, error) {
	path := "/customer/deactivate_authorization"

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}
