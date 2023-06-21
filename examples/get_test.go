package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/georgebent/go-httpclient/gohttp_mock"
)

func TestGetError(t *testing.T) {
	gohttp_mock.StartMockServer()

	gohttp_mock.AddMock(gohttp_mock.Mock{
		Method: http.MethodGet,
		Url:    "http://localhost",
		Error:  errors.New("timeout getting responce"),
	})

	result, err := Get()
	if result != "" {
		t.Error("Expect empty body")
	}

	if err == nil {
		t.Error("Expect not empty error")
	}

	if err.Error() != "timeout getting responce" {
		t.Error("Expect message 'timeout getting responce'")
	}
}

func TestGetResponse(t *testing.T) {
	gohttp_mock.AddMock(gohttp_mock.Mock{
		Method:         http.MethodGet,
		Url:            "http://localhost",
		ResponseStatus: 200,
		ResponseBody:   "Ok",
	})

	result, err := Get()
	if result == "" {
		t.Error("Expect not empty body")
	}

	if err != nil {
		t.Error("Expect empty error")
	}

	if result != "Ok" {
		t.Error("Expect message 'Ok'")
	}
}
