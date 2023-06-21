package examples

import (
	"fmt"
	"net/http"
)

func Post(url string) (string, error) {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-12345678")

	body := make(map[string]string)
	body["firstname"] = "John"
	body["lastname"] = "Stranger"
	body["type"] = "Builder Singletone"

	response, error := HttpClient.Post(url, headers, body)
	if error != nil {
		fmt.Println(string(error.Error()))

		return "", error
	}

	fmt.Println(response.StatusCode)

	return response.BodyString(), nil
}
