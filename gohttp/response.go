package gohttp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	status     string
	statusCode int
	headers    http.Header
	body       []byte
}

func (r *Response) GetStatus() string {
	return r.status
}

func (r *Response) GetStatusCode() int {
	return r.statusCode
}

func (r *Response) GetHeaders() http.Header {
	return r.headers
}

func (r *Response) BodyBytes() []byte {
	return r.body
}

func (r *Response) BodyString() string {
	return string(r.body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.BodyBytes(), target)
}
