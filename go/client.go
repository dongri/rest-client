package client

import (
	"fmt"
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

// ContentType ...
type ContentType string

// ContentType ...
const (
	ContentTypeMultipart  ContentType = "multipart-form-data"
	ContentTypeUrlencoded ContentType = "application/x-www-form-urlencoded"
	ContentTypeJSON       ContentType = "application/json"
	ContentTypeXML        ContentType = "application/xml"
	ContentTypeBase64     ContentType = "application/base64"
	ContentTypeStream     ContentType = "application/octet-stream"
	ContentTypePlain      ContentType = "text/plain"
	ContentTypeCss        ContentType = "text/css"
	ContentTypeHtml       ContentType = "text/html"
	ContentTypeJavascript ContentType = "application/javascript"
)

// Client Struct
type Client struct {
	BaseURI string
	Header  map[string][]string
	Timeout int
}

// NewClient ...
func NewClient(baseURI string, header map[string][]string, timeout int) *Client {
	return &Client{
		BaseURI: baseURI,
		Header:  header,
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
	switch {
	case method == GET || method == DELETE:
		req, err := http.NewRequest(string(method), c.BaseURI, nil)
		if err != nil {
			return nil, err
		}
		req.URL.Path = urlpath.Join(req.URL.Path, path)
		req.URL.RawQuery = vals.Encode()
		req.Header = c.Header
		return c.do(req)
	case method == POST || method == PUT:
		req, err := http.NewRequest(string(method), c.BaseURI, strings.NewReader(vals.Encode()))
		if err != nil {
			return nil, err
		}
		req.URL.Path = urlpath.Join(req.URL.Path, path)
		req.Header = c.Header
		return c.do(req)
	default:
		return nil, fmt.Errorf("Unsupport method: %v", method)
	}
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	httpClient := &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	return httpClient.Do(req)
}
