package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	urlpath "path"
	"strings"
	"time"
)

// ContentType ...
type ContentType string

// ContentType ...
const (
	ContentTypeFormUrlencoded ContentType = "application/x-www-form-urlencoded"
	ContentTypeJSON           ContentType = "application/json"
)

// Client Struct
type Client struct {
	BaseURI     string
	ContentType ContentType
	Header      map[string]string
	Timeout     int
}

// NewClient ...
func NewClient(baseURI string, contentType ContentType, header map[string]string, timeout int) *Client {
	return &Client{
		BaseURI:     baseURI,
		ContentType: contentType,
		Header:      header,
		Timeout:     timeout,
	}
}

// Get request GET.
func (c *Client) Get(path string, query map[string]string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodGet, path, query, params)
}

// Post request POST.
func (c *Client) Post(path string, query map[string]string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodPost, path, query, params)
}

// Put request PUT.
func (c *Client) Put(path string, query map[string]string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodPut, path, query, params)
}

// Delete request DELETE.
func (c *Client) Delete(path string, query map[string]string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodDelete, path, query, params)
}

// Patch request PATCH.
func (c *Client) Patch(path string, query map[string]string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodPatch, path, query, params)
}

func (c *Client) requestWithMethod(method string, path string, query map[string]string, params map[string]string) (*http.Response, error) {
	queries := url.Values{}
	for k, v := range query {
		queries.Add(k, v)
	}

	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v)
	}

	var reader io.Reader
	if c.ContentType == ContentTypeJSON {
		s, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer([]byte(s))
	} else if c.ContentType == ContentTypeFormUrlencoded {
		reader = strings.NewReader(vals.Encode())
	} else {
		return nil, errors.New("content-type error")
	}
	req, err := http.NewRequest(string(method), c.BaseURI, reader)
	if err != nil {
		return nil, err
	}

	req.URL.Path = urlpath.Join(req.URL.Path, path)
	req.URL.RawQuery = queries.Encode()

	req.Header.Set("Content-Type", string(c.ContentType))
	for k, v := range c.Header {
		req.Header.Set(k, v)
	}

	return c.do(req)
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	return httpClient.Do(req)
}
