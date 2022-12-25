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
	testCase := &TransactionBody{
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

func TestFetchTransaction(t *testing.T) {
	t.Run("gets details of a transactionn carried out on your integration", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		fetchTrnx, err := client.FetchTransaction(2)
		if err != nil {
			t.Error("cannot fetch transactions")
		}

		data, _ := json.MarshalIndent(fetchTrnx, "", "    ")
		fmt.Println(string(data))
	})
}

func TestChargeAuthorization(t *testing.T) {
	testCase := &ChargeAuthorizationBody{
		Amount:            "20",
		Email:             "test@test.com",
		AuthorizationCode: "a random auth code",
	}
	t.Run("charge authorization", func(t *testing.T) {
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

		if testCase.AuthorizationCode == "" {
			t.Error("please provide an authorization code")
		}

		response, err := client.ChargeAuthorization(testCase)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))

	})
}

func TestCheckAuthorization(t *testing.T) {
	testCase := &CheckAuthorizationBody{
		Amount:            "300",
		Email:             "test@test.mail",
		AuthorizationCode: "random code",
	}

	t.Run("check authorization", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		if testCase.Email == "" {
			t.Error("please provide an email")
		}

		if testCase.Amount == "" {
			t.Error("please provide an amount")
		}

		if testCase.AuthorizationCode == "" {
			t.Error("please provide an authorizationCode")
		}

		response, err := client.CheckAuthorization(testCase)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestViewTransactionTimeLine(t *testing.T) {
	t.Run("view transaction timeline", func(t *testing.T) {
		referenceId := "a random ref"
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.ViewTransactionTimeLine(referenceId)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestTransactionTotals(t *testing.T) {
	t.Run("transaction totals", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.TransactionTotals()
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestExportTransactions(t *testing.T) {
	t.Run("export transactions", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.ExportTransactions()
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestPartialDebit(t *testing.T) {
	testCase := &PartialDebit{
		AuthorizationCode: "random",
		Currency:          "NGN",
		Amount:            "20",
		Email:             "test@test.com",
	}

	t.Run("partial debit", func(t *testing.T) {
		client, err := NewClient(apiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.PartialDebit(testCase)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}
