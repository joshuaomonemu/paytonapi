package models

import (
	"io/ioutil"
	"net/http"
)

func AllCards() ([]byte, error) {

	req, _ := http.NewRequest("GET", url31, nil)

	req.Header.Add("Accept", "application/com.reloadly.giftcards-v1+json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
