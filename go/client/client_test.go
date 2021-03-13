package client

import (
	"encoding/json"
	"testing"
)

const (
	endpoint = "https://yourserver.herokuapp.com"
	// endpoint = "http://localhost:3000"
	timeout = 90 //Second
)

func TestGet(t *testing.T) {
	client := NewClient(endpoint, ContentTypeJSON, nil, timeout)
	query := map[string]string{
		"name": "dongri",
	}
	res, err := client.Get("/", query, nil)
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

func TestPostForm(t *testing.T) {
	header := map[string]string{
		"X-AccessToken": "hoge",
	}
	client := NewClient(endpoint, ContentTypeFormUrlencoded, header, timeout)
	params := map[string]string{
		"name": "dongri",
	}
	res, err := client.Post("/", nil, params)
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

func TestPostJSON(t *testing.T) {
	header := map[string]string{
		"X-AccessToken": "hoge",
	}
	client := NewClient(endpoint, ContentTypeJSON, header, timeout)
	params := map[string]string{
		"name": "dongri",
	}
	res, err := client.Post("/", nil, params)
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
	header := map[string]string{
		"X-AccessToken": "hoge",
	}
	client := NewClient(endpoint, ContentTypeFormUrlencoded, header, timeout)
	params := map[string]string{
		"name": "dongri",
	}
	res, err := client.Put("/", nil, params)
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
	client := NewClient(endpoint, ContentTypeJSON, nil, timeout)
	params := map[string]string{
		"name": "donri",
	}
	res, err := client.Delete("/", nil, params)
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
