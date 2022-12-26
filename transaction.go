// The Transactions API allows you create and manage payments on your integration

package paystack

import (
	"encoding/json"
	"fmt"
)

var (
	jsonResponse Response
)

type TransactionBody struct {
	// Amount should be in *kobo* if currency is *NGN*,
	// *pesewas*, if currency is *GHS*, and cents, if
	// currency is *ZAR*
	Amount string `json:"amount,omitempty"`

	// Email:  Customer email address
	Email string `json:"email,omitempty"`

	// Currency: The default currency (NGN, GHS, ZAR, or USD)
	// Defaults to your integration currency
	Currency string `json:"currency,omitempty"`

	// Reference: Unique transaction reference. Only -, ., =
	// and alphanumeric characters allowed.
	Reference string `json:"reference,omitempty"`

	// CallbackURL: Fully qualified url, e.g. *https://example.com*.
	// Use this to override the callback url provided on the dashboard for this transaction
	CallbackURL string `json:"callback_url,omitempty"`

	// Plan: If transaction is to create a subscription to a predefined plan,
	// provide plan code here. This would invalidate the value provided in amount
	Plan string `json:"plan,omitempty"`

	// InvoiceLimit: Number of times to charge customer during subscription to plan
	InvoiceLimit uint64 `json:"invoice_limit,omitempty"`

	// Metadata: Stringified JSON object of custom data.
	// Kindly check the Metadata page for more information.
	// https://paystack.com/docs/payments/metadata
	Metadata string `json:"metadata,omitempty"`

	// Channels: An array of payment channels to control what channels
	// you want to make available to the user to make a payment with.
	// Available channels include: ["card", "bank", "ussd", "qr", "mobile_money", "bank_transfer", "eft"]
	Channels []string `json:"channels,omitempty"`

	// SplitCode: The split code of the transaction split. e.g. SPL_98WF13Eb3w
	SplitCode string `json:"split_code,omitempty"`

	// Subaccount: The code for the subaccount that owns the payment. e.g. ACCT_8f4s1eq7ml6rlzj
	Subaccount string `json:"subaccount,omitempty"`

	// TransactionCharge: An amount used to override the split configuration
	// for a single split payment. If set, the amount specified goes
	// to the main account regardless of the split configuration.
	TransactionCharge uint64 `json:"transaction_charge,omitempty"`

	// Bearer: Who bears Paystack charges? account or subaccount (defaults to account).
	Bearer string `json:"bearer,omitempty"`
}

type ChargeAuthorizationBody struct {
	// Amount should be in kobo if currency is NGN, pesewas,
	// if currency is GHS, and cents, if currency is ZAR
	Amount string `json:"amount"`

	// Email: Customer's email address
	Email string `json:"email"`

	// AuthorizationCode: Valid authorization code to charge
	AuthorizationCode string `json:"authorization_code"`

	// Reference: Unique transaction reference. Only -, ., =
	// and alphanumeric characters allowed.
	Reference string `json:"reference,omitempty"`

	// Currency in which amount should be charged. Allowed values are: NGN, GHS, ZAR or USD
	Currency string `json:"currency,omitempty"`

	// Metadata: Stringified JSON object.
	// Add a custom_fields attribute which has an array of objects
	// if you would like the fields to be added to your transaction
	// when displayed on the dashboard.
	// Sample: {"custom_fields":[{"display_name":"Cart ID","variable_name": "cart_id","value": "8393"}]}
	Metadata string `json:"metadata,omitempty"`

	// Channels: Send us 'card' or 'bank' or 'card','bank' as
	// an array to specify what options to show the user paying
	Channels []string `json:"channels,omitempty"`

	// Subaccount: The code for the subaccount that owns the payment. e.g. ACCT_8f4s1eq7ml6rlzj
	Subaccount string `json:"subaccount,omitempty"`

	// TransactionCharge: A flat fee to charge the subaccount for this
	// transaction (in kobo if currency is NGN, pesewas, if currency is
	// GHS, and cents, if currency is ZAR). This overrides the split
	// percentage set when the subaccount was created.
	// Ideally, you will need to use this if you are splitting in flat
	// rates (since subaccount creation only allows for percentage split). e.g. 7000 for a 70 naira
	TransactionCharge uint64 `json:"transaction_charge,omitempty"`

	// Bearer: Who bears Paystack charges? account or subaccount (defaults to account).
	Bearer string `json:"bearer,omitempty"`

	// Queue: If you are making a scheduled charge call, it is a good idea
	// to queue them so the processing system does not get overloaded
	// causing transaction processing errors. Send queue:true to take advantage of our queued charging.
	Queue bool `json:"queue,omitempty"`
}

type CheckAuthorizationBody struct {
	// Amount should be in kobo if currency is NGN, pesewas,
	// if currency is GHS, and cents, if currency is ZAR
	Amount string `json:"amount"`

	// Email: Customer's email address
	Email string `json:"email"`

	// AuthorizationCode: Valid authorization code to charge
	AuthorizationCode string `json:"authorization_code"`

	// Currency in which amount should be charged. Allowed values are: NGN, GHS, ZAR or USD
	Currency string `json:"currency,omitempty"`
}

type PartialDebitBody struct {
	// Authorization Code
	AuthorizationCode string `json:"authorization_code"`

	// Currency: Specify the currency you want to debit. Allowed values are NGN, GHS, ZAR or USD.
	Currency string `json:"currency"`

	// Amount should be in kobo if currency is NGN, pesewas,
	// if currency is GHS, and cents, if currency is ZAR
	Amount string `json:"amount"`

	// Email: Customer's email address (attached to the authorization code)
	Email string `json:"email"`

	// Reference: Unique transaction reference. Only -, ., = and alphanumeric characters allowed.
	Reference string `json:"reference,omitempty"`

	// AtLeast: Minimum amount to charge
	AtLeast string `json:"at_least,omitempty"`
}

// InitializeTransaction initiate a new transaction
//
// Docs: https://paystack.com/docs/api/#transaction-initialize
//
//	client, _ := paystack.NewClient(apiKey)
//	transaction, err := client.InitializeTransaction(body struct{})
func (c *Config) InitializeTransaction(body *TransactionBody) (Response, error) {
	path := "/transaction/initialize"

	response, err := c.makeRequest("POST", path, body)
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
func (c *Config) VerifyTransaction(reference string) (Response, error) {
	path := fmt.Sprintf("/transaction/verify/%s", reference)
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

// CheckAuthorization:
// All Mastercard and Visa authorizations can be checked with this endpoint
// to know if they have funds for the payment you seek.
// This endpoint should be used when you do not know the exact amount to charge
// a card when rendering a service.
// It should be used to check if a card has enough funds based on a maximum range value. It is well suited for:
//
// *Feature Availability*: This feature is only available to businesses in Nigeria.
//
// *Warning*: You shouldn't use this endpoint to check a card for sufficient funds
// if you are going to charge the user immediately.
// This is because we hold funds when this endpoint is called which can lead to an insufficient funds error.
//
// Docs: https://paystack.com/docs/api/#transaction-check-authorization
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.CheckAuthorization(body *CheckAuthorizationBody struct{})
func (c *Config) CheckAuthorization(body *CheckAuthorizationBody) (Response, error) {
	path := "/transaction/check_authorization"
	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ViewTransactionTimeLine views the timeline of a transaction
//
// Docs: https://paystack.com/docs/api/#transaction-view-timeline
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ViewTransactionTimeLine(referenceOrID string)
func (c *Config) ViewTransactionTimeLine(referenceOrID string) (Response, error) {
	path := fmt.Sprintf("/transaction/timeline/%s", referenceOrID)
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// TransactionTotals returns the total amount received on your account
//
// Docs: https://paystack.com/docs/api/#transaction-totals
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.TransactionTotals()
func (c *Config) TransactionTotals() (Response, error) {
	path := "/transaction/totals"
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ExportTransactions lists out transactions carried out on your integration
// and export them to a csv file.
//
// Docs: https://paystack.com/docs/api/#transaction-export
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ExportTransactions()
func (c *Config) ExportTransactions() (Response, error) {
	path := "/transaction/export"
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// PartialDebit retrieves part of a payment from a customer
//
// Docs: https://paystack.com/docs/api/#transaction-partial-debit
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.PartialDebit(body struct{}})
func (c *Config) PartialDebit(body *PartialDebitBody) (Response, error) {
	path := "/transaction/partial_debit"
	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}
