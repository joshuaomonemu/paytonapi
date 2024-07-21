package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url21 = "https://giftcards-sandbox.reloadly.com/products"
)

func GetGiftCards() ([]byte, error) {
	if err2 != nil {
		return nil, err2
	}

	//payload := strings.NewReader(string(util))
	req, _ := http.NewRequest("GET", url21, nil)

	req.Header.Add("Accept", "application/com.reloadly.giftcards-v1+json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	fmt.Println(string(body))
	return body, nil
}
