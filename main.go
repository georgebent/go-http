package main

import (
	"fmt"
	"net/http"

	"github.com/georgebent/go-httpclient/gohttp"
)

func getAuthorisedClient() gohttp.HttpClient {
	httpClient := gohttp.New()

	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-1234")
	httpClient.SetHeaders(headers)

	return httpClient
}

func main() {
	client := getAuthorisedClient()
	client.DisableTimeouts(true)

	body := make(map[string]string)
	body["firstname"] = "John"
	body["lastname"] = "Stranger"

	response, error := client.Post("https://webhook.site/2c52f051-5e9f-458e-8e4d-4cf44fff1ada", nil, body)
	if error != nil {
		panic(error)
	}

	fmt.Println(response.StatusCode)
}
