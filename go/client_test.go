package client

import (
	"encoding/json"
	"fmt"
	"testing"
)

const timeout = 90 //second

func TestGet(t *testing.T) {
	client := NewClient("https://yourserver.herokuapp.com", timeout)
	params := map[string][]string{
		"page": {"1"},
	}
	res, err := client.Get("/users", params)
	if err != nil {
		t.Errorf("got error %v", err)
	}
	defer res.Body.Close()

	resBody := map[string]interface{}{}
	if err := json.NewDecoder(res.Body).Decode(resBody); err != nil {
		t.Errorf("got error %v", err)
	}
	fmt.Println(resBody["age"])
}
