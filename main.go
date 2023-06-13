package main

import (
	"fmt"

	"github.com/georgebent/go-httpclient/gohttp"
)

func main() {
	client := gohttp.New()
	response, error := client.Get("http://localhost", nil)
	if error != nil {
		panic(error)
	}

	fmt.Println(response.StatusCode)
}
