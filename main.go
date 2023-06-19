package main

import (
	"fmt"
	"net/http"

	"github.com/georgebent/go-httpclient/gohttp"
)

var (
	AuthorisedClient = getAuthorisedClient()
)

func getAuthorisedClient() gohttp.HttpClient {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-12345678")

	httpClient := gohttp.NewBuilder().SetHeaders(headers).DisableTimeouts(true).Build()

	return httpClient
}

func main() {
	runRequest()
}

func runRequest() {
	body := make(map[string]string)
	body["firstname"] = "John"
	body["lastname"] = "Stranger"
	body["type"] = "Builder Singletone"

	response, error := AuthorisedClient.Post("https://webhook.site/2c52f051-5e9f-458e-8e4d-4cf44fff1ada", nil, body)
	if error != nil {
		panic(error)
	}

	fmt.Println(response.GetStatusCode())
}
