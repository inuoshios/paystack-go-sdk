package paystack

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateSplit(t *testing.T) {
	testCase := &CreateSplitBody{}

	t.Run("create split", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.CreateSplit(testCase)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestListAndSearchSplits(t *testing.T) {
	t.Run("list and search splits", func(t *testing.T) {
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.ListAndSearchSplits()
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}

func TestFetchSplit(t *testing.T) {
	t.Run("fetch split", func(t *testing.T) {
		query := "143"
		client, err := NewClient(ApiKey)
		if err != nil {
			t.Errorf("cannot initialize client %s", err)
		}

		response, err := client.FetchSplit(query)
		if err != nil {
			t.Error(err)
		}

		data, _ := json.MarshalIndent(response, "", "    ")
		fmt.Println(string(data))
	})
}
