package gohttp_mock

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type httpClientMock struct {
}

func (c *httpClientMock) Do(request *http.Request) (*http.Response, error) {
	requestBody, error := request.GetBody()
	if error != nil {
		return nil, error
	}

	defer requestBody.Close()

	body, error := io.ReadAll(requestBody)
	if error != nil {
		return nil, error
	}

	var response http.Response

	mock := MockupServer.mocks[MockupServer.getMockKey(request.Method, request.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		response.StatusCode = mock.ResponseStatus
		response.Body = io.NopCloser(strings.NewReader(mock.ResponseBody))
		response.ContentLength = int64(len(mock.ResponseBody))
		response.Request = request
		return &response, nil
	}

	return nil, fmt.Errorf("no mock matching %s from '%s' with given body", request.Method, request.URL.String())
}
