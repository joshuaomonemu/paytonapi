package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	data  = "https://vtpass.com/api/service-variations?serviceID="
	data1 = "https://vtpass.com/api/pay"
)

func Data(network string) ([]byte, error) {
	req, _ := http.NewRequest("GET", data+network, nil)

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

func DataPay(biller, provider, amount, phone, variation_code, request_id string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(data1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, err
	}

	// Create query parameters
	params := url.Values{}
	params.Add("billersCode", biller)
	params.Add("serviceID", provider)
	params.Add("amount", amount)
	params.Add("phone", phone)
	params.Add("variation_code", variation_code)
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
