package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.github.com"

	client := http.Client{}

	response, error := client.Get(url)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()
	bytes, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
