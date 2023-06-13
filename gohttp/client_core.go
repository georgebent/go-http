package gohttp

import "net/http"

func (c *Client) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, error := http.NewRequest(method, url, nil)
	if error != nil {
		return nil, error
	}

	return client.Do(request)
}
