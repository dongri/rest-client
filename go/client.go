package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
func (c *Client) Get(path string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodGet, path, params)
}

// Post request POST.
func (c *Client) Post(path string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodPost, path, params)
}

// Put request PUT.
func (c *Client) Put(path string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodPut, path, params)
}

// Delete request DELETE.
func (c *Client) Delete(path string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodDelete, path, params)
}

// Patch request PATCH.
func (c *Client) Patch(path string, params map[string]string) (*http.Response, error) {
	return c.requestWithMethod(http.MethodPatch, path, params)
}

func (c *Client) requestWithMethod(method string, path string, params map[string]string) (*http.Response, error) {
	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v)
	}
	switch {
	case method == http.MethodGet || method == http.MethodDelete:
		req, err := http.NewRequest(string(method), c.BaseURI, nil)
		if err != nil {
			return nil, err
		}
		req.URL.Path = urlpath.Join(req.URL.Path, path)
		req.URL.RawQuery = vals.Encode()
		req.Header.Set("Content-Type", string(c.ContentType))
		for k, v := range c.Header {
			req.Header.Set(k, v)
		}
		return c.do(req)
	case method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch:
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
		req.Header.Set("Content-Type", string(c.ContentType))
		for k, v := range c.Header {
			req.Header.Set(k, v)
		}
		return c.do(req)
	default:
		return nil, fmt.Errorf("Unsupport method: %v", method)
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	return httpClient.Do(req)
}
