package examples

func Get() (string, error) {
	res, err := HttpClient.Get("http://localhost", nil)
	if err != nil {
		return "", err
	}

	return res.BodyString(), nil
}
