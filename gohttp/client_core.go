package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	DEFAULT_MAX_IDDLE_CONNECTIONS = 5
	DEFAULT_RESPONSE_TIMEOUT      = 5 * time.Second
	DEFAULT_CONNECTION_TIMEOUT    = 1 * time.Second
)

func (c *Client) do(method string, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.getRequestHeaders(headers)
	requestBody, error := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if error != nil {
		return nil, error
	}

	if mock := mockupServer.getMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}

	request, error := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if error != nil {
		return nil, error
	}

	request.Header = fullHeaders

	response, error := c.getHttpClient().Do(request)
	if error != nil {
		return nil, error
	}

	defer response.Body.Close()
	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}

	finalResponse := Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		headers:    request.Header,
		body:       responseBody,
	}

	finalResponse.status = "dsd ssd"

	return &finalResponse, nil
}

func (c *Client) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.Builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	return result
}

func (c *Client) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *Client) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		c.CoreClient = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConns:          c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.CoreClient
}

func (c *Client) getMaxIdleConnections() int {
	if c.Builder.maxIdleConnections > 0 {
		return c.Builder.maxIdleConnections
	}

	return DEFAULT_MAX_IDDLE_CONNECTIONS
}

func (c *Client) getResponseTimeout() time.Duration {
	if c.Builder.responseTimeout > 0 {
		return c.Builder.responseTimeout
	}

	if c.Builder.disabledTimeouts {
		return 0
	}

	return DEFAULT_RESPONSE_TIMEOUT
}

func (c *Client) getConnectionTimeout() time.Duration {
	if c.Builder.connectionTimeout > 0 {
		return c.Builder.connectionTimeout
	}

	if c.Builder.disabledTimeouts {
		return 0
	}

	return DEFAULT_CONNECTION_TIMEOUT
}
