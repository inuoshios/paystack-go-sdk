package paystack

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	apiKey = "sk_test_5dc4a0ff21ab4f7287efbc0951794af82a6a2b48"
)

func TestInitializeTransaction(t *testing.T) {
	testCase := &InitTrnx{
		Amount:   "300",
		Email:    "test@test.com",
		Currency: "NGN",
	}

	t.Run("initialize a new transaction", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		if testCase.Amount == "" {
			t.Error("please provide an amount")
		}

		if testCase.Email == "" {
			t.Error("please provide an email")
		}

		response, err := client.InitializeTransaction(testCase)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestVerifyTransaction(t *testing.T) {
	transactionReference := "dm9jdrejvp"

	t.Run("verify a transaction using the transaction reference", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		verifyTransaction, err := client.VerifyTransaction(transactionReference)
		if err != nil {
			t.Error("cannot verify transaction")
		}

		data, _ := json.MarshalIndent(verifyTransaction, "", "    ")
		fmt.Println(string(data))
	})
}

func TestListTransaction(t *testing.T) {
	t.Run("list transaction", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		listTrnx, err := client.ListTransactions()
		if err != nil {
			t.Error("cannot list transaction")
		}

		data, _ := json.MarshalIndent(listTrnx, "", "    ")
		fmt.Println(string(data))
	})
}
