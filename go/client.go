package client

import (
	"net/http"
	"net/url"
	urlpath "path"
	"strings"
	"time"
)

// HTTPMethod ...
type HTTPMethod string

// HTTPMethod ...
const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

// Client Struct
type Client struct {
	BaseURI string
	Timeout int
}

// NewClient ...
func NewClient(baseURI string, timeout int) *Client {
	return &Client{
		BaseURI: baseURI,
		Timeout: timeout,
	}
}

// Get request GET.
func (c *Client) Get(path string, params map[string][]string) (*http.Response, error) {
	return c.requestWithMethod(GET, path, params)
}

// Post request POST.
func (c *Client) Post(path string, params map[string][]string) (*http.Response, error) {
	return c.requestWithMethod(POST, path, params)
}

// Put request PUT.
func (c *Client) Put(path string, params map[string][]string) (*http.Response, error) {
	return c.requestWithMethod(PUT, path, params)
}

// Delete request DELETE.
func (c *Client) Delete(path string, params map[string][]string) (*http.Response, error) {
	return c.requestWithMethod(DELETE, path, params)
}

// Patch request PATCH.
func (c *Client) Patch(path string, params map[string][]string) (*http.Response, error) {
	return c.requestWithMethod(PATCH, path, params)
}

func (c *Client) requestWithMethod(method HTTPMethod, path string, params map[string][]string) (*http.Response, error) {
	vals := url.Values(params)
	if method == GET {
		req, err := http.NewRequest(string(GET), c.BaseURI, nil)
		if err != nil {
			return nil, err
		}
		req.URL.Path = urlpath.Join(req.URL.Path, path)
		req.URL.RawQuery = vals.Encode()
		return c.do(req)
	}
	req, err := http.NewRequest(string(method), c.BaseURI, strings.NewReader(vals.Encode()))
	if err != nil {
		return nil, err
	}
	req.URL.Path = urlpath.Join(req.URL.Path, path)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.do(req)
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	return httpClient.Do(req)
}
