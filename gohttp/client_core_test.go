package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	client := Client{}

	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid Content-Type received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("Invalid User-Agent received")
	}

	if finalHeaders.Get("X-Request-Id") != "123" {
		t.Error("Invalid X-Request-Id received")
	}
}

func TestGetRequestBody(t *testing.T) {
	client := Client{}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, error := client.getRequestBody("", nil)

		if error != nil {
			t.Error("Expected nil error when passing a nil body")
		}

		if body != nil {
			t.Error("Expected nil body when passing a nil body")
		}
	})

	t.Run("BodyJsonResponse", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, error := client.getRequestBody("application/json", requestBody)
		if error != nil {
			t.Error("Expected nil error when passing a json body")
		}

		if string(body) != `["one","two"]` {
			t.Error("Wrong body when passing a json body")
		}
	})

	t.Run("BodyXmlResponse", func(t *testing.T) {
		requestBody := []string{"one", "two"}

		body, error := client.getRequestBody("application/xml", requestBody)
		if error != nil {
			t.Error("Expected nil error when passing a json body")
		}

		if string(body) != "<string>one</string><string>two</string>" {
			t.Error("Wrong body when passing a json body")
		}
	})

	t.Run("BodyDefaultResponse", func(t *testing.T) {
		requestBody := []string{"three", "four"}

		body, error := client.getRequestBody("", requestBody)
		if error != nil {
			t.Error("Expected nil error when passing a json body")
		}

		if string(body) != `["three","four"]` {
			t.Error("Wrong body when passing a json body")
		}
	})
}
