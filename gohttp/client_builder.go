package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	headers            http.Header
	disabledTimeouts   bool
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
	c.headers = headers

	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout

	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout

	return c
}

func (c *clientBuilder) SetMaxIdleConnections(maxConnections int) ClientBuilder {
	c.maxIdleConnections = maxConnections

	return c
}

func (c *clientBuilder) DisableTimeouts(disabledTimeouts bool) ClientBuilder {
	c.disabledTimeouts = disabledTimeouts

	return c
}
