package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

func (c *Client) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
	fullHeaders := c.getRequestHeaders(headers)
	requestBody, error := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if error != nil {
		return nil, error
	}

	request, error := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if error != nil {
		return nil, error
	}

	request.Header = fullHeaders

	return client.Do(request)
}

func (c *Client) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	for header, value := range c.Headers {
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
