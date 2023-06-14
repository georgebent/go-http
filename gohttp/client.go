package gohttp

import (
	"net/http"
)

type HttpClient interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
}

func New() HttpClient {
	client := &Client{}
	return client
}

type Client struct {
	Headers http.Header
}

func (c *Client) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *Client) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *Client) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *Client) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, nil)
}

func (c *Client) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
func (c *Client) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}
