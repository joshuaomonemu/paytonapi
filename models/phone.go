package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	phoneurl = "https://vtpass.com/api/pay"
)

func PhonePay(provider, amount, phone, request_id string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(phoneurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, err
	}

	// Create query parameters
	params := url.Values{}
	params.Add("serviceID", provider)
	params.Add("amount", amount)
	params.Add("phone", phone)
	params.Add("request_id", request_id)

	// Add the query parameters to the URL
	u.RawQuery = params.Encode()

	req, _ := http.NewRequest("POST", u.String(), nil)

	req.Header.Add("api-key", api)
	req.Header.Add("public-key", public)
	req.Header.Add("secret-key", secret)

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
