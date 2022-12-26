package paystack

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreatePlan(t *testing.T) {
	createPlan := &Plan{
		Name:        "Montly retainer",
		Amount:      300,
		Interval:    "monthly",
		Description: "Monthly plan",
		Currency:    "NGN",
	}

	t.Run("create plan", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.CreatePlan(createPlan)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestListPlans(t *testing.T) {
	t.Run("list plans", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.ListPlans()
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestFetchPlan(t *testing.T) {
	t.Run("test plan", func(t *testing.T) {
		codeOrID := "PLN_gx2wn530m0i3w3m"

		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.FetchPlan(codeOrID)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestUpdatePlan(t *testing.T) {
	updatePlan := &Plan{
		Name: "Montly retainer",
	}

	t.Run("create plan", func(t *testing.T) {
		codeOrID := "PLN_gx2wn530m0i3w3m"

		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.UpdatePlan(codeOrID, updatePlan)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}
