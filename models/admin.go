package models

import (
	"io/ioutil"
	"net/http"
)

var (
	bala = "https://vtpass.com/api/balance"
)

func Balance() ([]byte, error) {
	req, _ := http.NewRequest("GET", bala, nil)

	req.Header.Add("api-key", api)
	req.Header.Add("public-key", public)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
