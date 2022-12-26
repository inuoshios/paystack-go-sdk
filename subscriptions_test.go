package paystack

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateSubscription(t *testing.T) {
	createSubscription := &CreateSubscriptionBody{
		Customer: "",
		Plan:     "",
	}

	t.Run("create subscription", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		createSub, err := client.CreateSubscription(createSubscription)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(createSub, "", "    ")
		fmt.Println(string(data))
	})
}

func TestListSubscriptions(t *testing.T) {
	t.Run("list subscriptions", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		listSub, err := client.ListSubscriptions()
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(listSub, "", "    ")
		fmt.Println(string(data))
	})
}

func TestFetchSubscription(t *testing.T) {
	t.Run("fetch subscription", func(t *testing.T) {
		codeOrID := "random"

		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		listSub, err := client.FetchSubscription(codeOrID)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(listSub, "", "    ")
		fmt.Println(string(data))

	})
}

func TestEnableSubscription(t *testing.T) {
	enableSub := &SubscriptionBody{
		Code:  "",
		Token: "",
	}

	t.Run("enable subscription", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		listSub, err := client.EnableSubscription(enableSub)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(listSub, "", "    ")
		fmt.Println(string(data))
	})
}

func TestDisableSubscription(t *testing.T) {
	disableSub := &SubscriptionBody{
		Code:  "",
		Token: "",
	}

	t.Run("disable subscription", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		listSub, err := client.DisableSubscription(disableSub)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(listSub, "", "    ")
		fmt.Println(string(data))
	})
}

func TestSendUpdateSubscriptionLink(t *testing.T) {
	t.Run("send update subscription link", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		sub, err := client.SendUpdateSubscriptionLink("random")
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(sub, "", "    ")
		fmt.Println(string(data))
	})
}

func TestGenerateUpdateSubscriptionLink(t *testing.T) {
	t.Run("generate update subscription link", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		sub, err := client.GenerateUpdateSubscriptionLink("random")
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(sub, "", "    ")
		fmt.Println(string(data))
	})
}
