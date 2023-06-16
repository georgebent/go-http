package gohttp

import (
	"net/http"
	"time"
)

type HttpClient interface {
	SetHeaders(headers http.Header)
	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	SetConnectionTimeout(timeout time.Duration)
	SetResponseTimeout(timeout time.Duration)
	SetMaxIdleConnections(maxConnections int)
	DisableTimeouts(b bool)
}

type Client struct {
	CoreClient         *http.Client
	MaxIdleConnections int
	ConnectionTimeout  time.Duration
	ResponseTimeout    time.Duration
	Headers            http.Header
	DisabledTimeouts   bool
}

func New() HttpClient {
	client := &Client{}

	return client
}

func (c *Client) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *Client) SetConnectionTimeout(timeout time.Duration) {
	c.ConnectionTimeout = timeout
}

func (c *Client) SetResponseTimeout(timeout time.Duration) {
	c.ResponseTimeout = timeout
}

func (c *Client) SetMaxIdleConnections(maxConnections int) {
	c.MaxIdleConnections = maxConnections
}

func (c *Client) DisableTimeouts(disabledTimeouts bool) {
	c.DisabledTimeouts = disabledTimeouts
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
