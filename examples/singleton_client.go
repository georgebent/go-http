package examples

import (
	"net/http"
	"time"

	"github.com/georgebent/go-httpclient/gohttp"
	"github.com/georgebent/go-httpclient/gomime"
)

var (
	HttpClient = getHttpClient()
)

func getHttpClient() gohttp.HttpClient {
	headers := make(http.Header)
	headers.Set(gomime.HEADER_CONTENT_TYPE, gomime.CONTENT_TYPE_JSON)

	client := gohttp.
		NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(3 * time.Second).
		SetConnectionTimeout(3 * time.Second).
		SetUserAgent("desktop").
		Build()

	return client
}
