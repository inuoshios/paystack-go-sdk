package paystack

import (
	"net/http"
)

var (
// baseUrl   = "https://api.paystack.co/"
)

func NewClient(apiKey string) *http.Header {
	return &http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{apiKey},
	}
}
