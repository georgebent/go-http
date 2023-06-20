package examples

import (
	"errors"
	"net/http"
	"testing"

	"github.com/georgebent/go-httpclient/gohttp"
)

func TestGetError(t *testing.T) {
	gohttp.StartMockServer()

	gohttp.AddMock(gohttp.Mock{
		Method: http.MethodGet,
		Url:    "https://webhook.site/2c52f051-5e9f-458e-8e4d-4cf44fff1ada",
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
	gohttp.AddMock(gohttp.Mock{
		Method:         http.MethodGet,
		Url:            "https://webhook.site/2c52f051-5e9f-458e-8e4d-4cf44fff1ada",
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
