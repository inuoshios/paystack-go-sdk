// The Transaction Splits API enables merchants split the
// settlement for a transaction across their payout account, and one or more Subaccounts.
package paystack

import (
	"encoding/json"
	"fmt"
)

type CreateSplit struct {
	// Name of the transaction split
	Name string `json:"name"`

	// Type: The type of transaction split you want to create.
	// You can use one of the following: percentage | flat
	Type string `json:"type"`

	// Currency: Any of NGN, GHS, ZAR, or USD
	Currency string `json:"currency"`

	// Subaccounts: A list of object containing subaccount code
	// and number of shares: [{subaccount: ‘ACT_xxxxxxxxxx’, share: xxx},{...}]
	Subaccounts []map[string]any

	// BearerType: Any of subaccount | account | all-proportional | all
	BearerType string `json:"bearer_type"`

	// BearerSubAccount: Subaccount code
	BearerSubAccount string `json:"bearer_subaccount"`
}

type UpdateSplit struct {
	// Name of the transaction split
	Name string `json:"name"`

	// Active: True or False
	Active bool `json:"active"`

	// Any of the following values: subaccount | account | all-proportional | all
	BearerType string `json:"bearer_type"`

	// BearerSubAccount: Subaccount code of a subaccount in the split group.
	// This should be specified only if the bearer_type is subaccount
	BearerSubAccount string `json:"bearer_subaccount"`
}

type AddAndUpdateSplitSubaccount struct {
	// Subaccount: This is the sub account code
	Subaccount string `json:"subaccount"`

	// Share: This is the transaction share for the subaccount
	Share uint64 `json:"share"`
}

type RemoveSubAccountFromSplit struct {
	// Subaccount This is the sub account code
	Subaccount string `json:"subaccount"`
}

// CreateSplit create a split payment on your integration
//
// Docs: https://paystack.com/docs/api/#split-create
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.CreateSplit(body struct{}})
func (c *Config) CreateSplit(body *CreateSplit) (Response, error) {
	path := "/split"
	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// ListAndSearchSplits: list/search for the transaction splits available on your integration.
//
// Docs: https://paystack.com/docs/api/#split-list
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.ListAndSearchSplits()
func (c *Config) ListAndSearchSplits() (Response, error) {
	path := "/split"
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// FetchSplit: Get details of a split on your integration.

// Docs: https://paystack.com/docs/api/#split-fetch
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.FetchSplit(query string)
func (c *Config) FetchSplit(query string) (Response, error) {
	path := fmt.Sprintf("/split/%s", query)
	response, err := c.makeRequest("GET", path, nil)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}

	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// UpdateSplit: Update a transaction split details on your integration
//
// Docs: https://paystack.com/docs/api/#split-update
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.UpdateSplit(body struct{}, query string)
func (c *Config) UpdateSplit(body *UpdateSplit, query string) (Response, error) {
	path := fmt.Sprintf("/split/%s", query)
	response, err := c.makeRequest("PUT", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// AddAndUpdateSplitSubaccount: Add a Subaccount to a Transaction Split,
// or update the share of an existing Subaccount in a Transaction Split
//
// Docs: https://paystack.com/docs/api/#split-add-subaccount
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.AddAndUpdateSplitSubaccount(body struct{}, query string)
func (c *Config) AddAndUpdateSplitSubaccount(body *AddAndUpdateSplitSubaccount, query string) (Response, error) {
	path := fmt.Sprintf("/split/%s/subaccount/add", query)
	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}

// RemoveSubAccountFromSplit: Remove a subaccount from a transaction split
//
// Docs: https://paystack.com/docs/api/#split-remove-subaccount
//
//	client, _ := paystack.NewClient(apiKey)
//	auth, err := client.RemoveSubAccountFromSplit(body struct{}, query string)
func (c *Config) RemoveSubAccountFromSplit(body *RemoveSubAccountFromSplit, query string) (Response, error) {
	path := fmt.Sprintf("/split/%s/subaccount/remove", query)
	response, err := c.makeRequest("POST", path, body)
	if err != nil {
		_ = json.Unmarshal(response, &jsonResponse)
		return jsonResponse, err
	}
	_ = json.Unmarshal(response, &jsonResponse)
	return jsonResponse, err
}
