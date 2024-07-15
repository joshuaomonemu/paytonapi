package models

import (
	"app/auth"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	url11       = "https://utilities-sandbox.reloadly.com/billers"
	url12       = "https://utilities-sandbox.reloadly.com/pay"
	token, err2 = auth.Auth2()
)

func LoadBillers() ([]byte, error) {
	if err2 != nil {
		return nil, err2
	}

	req, _ := http.NewRequest("GET", url11, nil)

	req.Header.Add("Accept", "application/com.reloadly.utilities-v1+json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func PayBill(util []byte) ([]byte, error) {
	if err2 != nil {
		return nil, err2
	}

	payload := strings.NewReader(string(util))

	req, _ := http.NewRequest("POST", url12, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/com.reloadly.utilities-v1+json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// fmt.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
