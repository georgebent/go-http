package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost"

	client := http.Client{}

	response, error := client.Get(url)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	bytes, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
