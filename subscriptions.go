// The Subscriptions API allows you create and manage recurring payment on your integration

package paystack

import (
	"encoding/json"
	"fmt"
)

type CreateSubscriptionBody struct {
	// Customer's email address or customer code
	Customer string `json:"customer"`

	// Plan code
	Plan string `json:"plan"`

	// If customer has multiple authorizations, you can set the
	// desired authorization you wish to use for this subscription here.
	// If this is not supplied, the customer's most recent authorization would be used
	Authorization string `json:"authorization"`

	// Set the date for the first debit. (ISO 8601 format) e.g. 2017-05-16T00:30:13+01:00
	StartDate string `json:"start_date"`
}

type SubscriptionBody struct {
	// Subscription code
	Code string `json:"code"`

	// Email token
	Token string `json:"token"`
}

// CreateSubscription creates a subscription on your integration.
//
// Docs: https://paystack.com/docs/api/#subscription-create
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.CreateSubscription(body structs{})
func (c *Config) CreateSubscription(body *CreateSubscriptionBody) (Response, error) {
	path := "/subscription"

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ListSubscriptions list subscription available on your integration.
//
// Docs: https://paystack.com/docs/api/#subscription-list
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ListSubscriptions()
func (c *Config) ListSubscriptions() (Response, error) {
	path := "/subscription"

	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// FetchSubscription get details of a subscription on your integration.
//
// Docs: https://paystack.com/docs/api/#subscription-fetch
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.FetchSubscription(codeOrID string)
func (c *Config) FetchSubscription(codeOrID string) (Response, error) {
	path := fmt.Sprintf("/subscription/%s", codeOrID)

	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// EnableSubscription enables a subscription on your integration
//
// Docs: https://paystack.com/docs/api/#subscription-enable
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.EnableSubscription(body struct{})
func (c *Config) EnableSubscription(body *SubscriptionBody) (Response, error) {
	path := "/subscription/enable"

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// DisableSubscription disables a subscription on your integration
//
// Docs: https://paystack.com/docs/api/#subscription-disable
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.DisableSubscription(body struct{})
func (c *Config) DisableSubscription(body *SubscriptionBody) (Response, error) {
	path := "/subscription/enable"

	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// GenerateUpdateSubscriptionLink generates a link for updating the card on a subscription
//
// Docs: https://paystack.com/docs/api/#subscription-manage-link
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.GenerateUpdateSubscriptionLink(code string})
func (c *Config) GenerateUpdateSubscriptionLink(code string) (Response, error) {
	path := fmt.Sprintf("/subscription/%s/manage/link/", code)

	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// SendUpdateSubscriptionLink emails a customer a link for updating the card on their subscription
//
// Docs: https://paystack.com/docs/api/#subscription-manage-email
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.SendUpdateSubscriptionLink(code string})
func (c *Config) SendUpdateSubscriptionLink(code string) (Response, error) {
	path := fmt.Sprintf("/subscription/%s/manage/email/", code)

	response, err := c.makeRequest("POST", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}
