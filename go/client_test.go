package client

import (
	"encoding/json"
	"testing"
)

const timeout = 90 //second
const endpoint = "http://localhost:8080"

func TestGet(t *testing.T) {
	client := NewClient(endpoint, timeout)
	params := map[string][]string{
		"page": {"1"},
	}
	res, err := client.Get("/users", params)
	if err != nil {
		t.Errorf("got error %v", err)
	}
	defer res.Body.Close()

	var resBody interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		t.Errorf("got error %v", err)
	}
	t.Log(resBody)
}

func TestPost(t *testing.T) {
	client := NewClient(endpoint, timeout)
	params := map[string][]string{
		"name":  {"dognri"},
		"email": {"dognri@domain.com"},
	}
	res, err := client.Post("/users", params)
	if err != nil {
		t.Errorf("got error %v", err)
	}
	defer res.Body.Close()
	var resBody interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		t.Errorf("got error %v", err)
	}
	t.Log(resBody)
}

func TestPut(t *testing.T) {
	client := NewClient(endpoint, timeout)
	params := map[string][]string{
		"name":  {"dognri"},
		"email": {"dognri@domain.com"},
	}
	res, err := client.Put("/users", params)
	if err != nil {
		t.Errorf("got error %v", err)
	}
	defer res.Body.Close()
	var resBody interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		t.Errorf("got error %v", err)
	}
	t.Log(resBody)
}

func TestDelete(t *testing.T) {
	client := NewClient(endpoint, timeout)
	params := map[string][]string{}
	res, err := client.Delete("/users/1", params)
	if err != nil {
		t.Errorf("got error %v", err)
	}
	defer res.Body.Close()
	var resBody interface{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		t.Errorf("got error %v", err)
	}
	t.Log(resBody)
}
