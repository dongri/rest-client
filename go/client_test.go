package client

import (
	"encoding/json"
	"testing"
)

const timeout = 90 //second
const endpoint = "http://localhost:8080"

func TestGet(t *testing.T) {
	client := NewClient(endpoint, nil, timeout)
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
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	client := NewClient(endpoint, header, timeout)
	params := map[string][]string{
		"name":  {"dongri"},
		"email": {"dongri@domain.com"},
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
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	client := NewClient(endpoint, header, timeout)
	params := map[string][]string{
		"name":  {"dongri"},
		"email": {"dongri@domain.com"},
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
	client := NewClient(endpoint, nil, timeout)
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
