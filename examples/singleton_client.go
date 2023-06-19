package examples

import (
	"time"

	"github.com/georgebent/go-httpclient/gohttp"
)

var (
	HttpClient = getHttpClient()
)

func getHttpClient() gohttp.HttpClient {
	client := gohttp.
		NewBuilder().
		SetConnectionTimeout(3 * time.Second).
		SetConnectionTimeout(3 * time.Second).
		Build()

	return client
}
