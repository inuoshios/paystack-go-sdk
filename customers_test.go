package paystack

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateCustomer(t *testing.T) {
	createCustomer := &CreateCustomerBody{
		Email:     "test@test.com",
		FirstName: "Test",
		LastName:  "Test2",
		Phone:     "+2348123456789",
	}

	t.Run("create customers", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		if createCustomer.Email == "" {
			t.Error("please provide an email")
		}

		response, err := client.CreateCustomer(createCustomer)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestListCustomers(t *testing.T) {
	t.Run("list customers", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.ListCustomers()
		if err != nil {
			t.Error(err)
		}
		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestFetchCustomer(t *testing.T) {
	t.Run("fetch customers", func(t *testing.T) {
		customerEmail := "test@test.com"
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.FetchCustomer(customerEmail)
		if err != nil {
			t.Error(err)
		}
		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestUpdateCustomer(t *testing.T) {
	updateCustomer := &UpdateCustomerBody{
		FirstName: "Tester",
		LastName:  "Tester2",
	}

	t.Run("update customer", func(t *testing.T) {
		code := "CUS_42qtajqgknfkqgy"
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.UpdateCustomer(code, updateCustomer)
		if err != nil {
			t.Error(err)
		}
		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestValidateCustomer(t *testing.T) {
	validateCustomer := &ValidateCustomerBody{
		FirstName:     "Tester",
		LastName:      "Tester2",
		AccountNumber: "0123456789",
		BVN:           "20012345677",
		BankCode:      "007",
		Country:       "NG",
		Type:          "bank_account",
	}

	t.Run("validate customers", func(t *testing.T) {
		code := "CUS_42qtajqgknfkqgy"
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.ValidateCustomer(code, validateCustomer)
		if err != nil {
			t.Error(err)
		}
		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestWhiteListOrBlacklistCustomer(t *testing.T) {
	testCase := &WhiteListOrBlacklistCustomerBody{
		Customer:   "test@test.com",
		RiskAction: "allow",
	}
	t.Run("whitelist or blacklist customers", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		if testCase.Customer == "" {
			t.Error("please provide an email")
		}

		response, err := client.WhiteListOrBlacklistCustomer(testCase)
		if err != nil {
			t.Error(err)
		}
		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestDeactivateAuthorization(t *testing.T) {
	deactivate := &DeactivateAuthorizationBody{
		AuthorizationCode: "AUTH_72btv547",
	}

	t.Run("deactivate authorization", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}
		if deactivate.AuthorizationCode == "" {
			t.Error("please provide an authorization code")
		}

		response, err := client.DeactivateAuthorization(deactivate)
		if err != nil {
			t.Error(err)
		}
		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}
