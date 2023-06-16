package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	MaxIdleConnections int
	ConnectionTimeout  time.Duration
	ResponseTimeout    time.Duration
	Headers            http.Header
	DisabledTimeouts   bool
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(maxConnections int) ClientBuilder
	DisableTimeouts(b bool) ClientBuilder
	Build() HttpClient
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}

	return builder
}

func (c *clientBuilder) Build() HttpClient {
	client := Client{
		Builder: c,
	}

	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.Headers = headers

	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.ConnectionTimeout = timeout

	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.ResponseTimeout = timeout

	return c
}

func (c *clientBuilder) SetMaxIdleConnections(maxConnections int) ClientBuilder {
	c.MaxIdleConnections = maxConnections

	return c
}

func (c *clientBuilder) DisableTimeouts(disabledTimeouts bool) ClientBuilder {
	c.DisabledTimeouts = disabledTimeouts

	return c
}
