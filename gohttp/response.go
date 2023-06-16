package gohttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Response) GetStatus() string {
	return r.Status
}

func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

func (r *Response) GetHeaders() http.Header {
	return r.Headers
}

func (r *Response) BodyBytes() []byte {
	return r.Body
}

func (r *Response) BodyString() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.BodyBytes(), target)
}
