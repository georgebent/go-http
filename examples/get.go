package examples

import "fmt"

func Get() (string, error) {
	res, err := HttpClient.Get("https://webhook.site/2c52f051-5e9f-458e-8e4d-4cf44fff1ada", nil)
	if err != nil {
		return "", err
	}

	fmt.Println(res.BodyString())
	fmt.Println(res.GetStatusCode())

	return res.BodyString(), nil
}
