package gohttp

import (
	"fmt"
	"net/http"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string
	Error       error

	ResponseBody   string
	ResponseStatus int
}

func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		statusCode: m.ResponseStatus,
		body:       []byte(m.ResponseBody),
		status:     fmt.Sprintf("%d %s", m.ResponseStatus, http.StatusText(m.ResponseStatus)),
	}

	return &response, nil
}
