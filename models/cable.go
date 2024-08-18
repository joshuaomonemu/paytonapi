package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	dstv1  = "https://sandbox.vtpass.com/api/service-variations?serviceID=dstv"
	dstv2  = "https://sandbox.vtpass.com/api/merchant-verify"
	dstv3  = "https://sandbox.vtpass.com/api/pay"
	gotv1  = "https://sandbox.vtpass.com/api/service-variations?serviceID=gotv"
	gotv2  = "https://sandbox.vtpass.com/api/merchant-verify"
	gotv3  = "https://sandbox.vtpass.com/api/pay"
	api    = "6106a6abec1114cc73e59c24b634d6c4"
	public = "PK_3848b4365f111a6bd186fcf9be401c4da86d92d9a5a"
	secret = "SK_514a719059bf25078427bd53a23499b5f2f4eab5aaf"
)

func Dstv() ([]byte, error) {
	req, _ := http.NewRequest("GET", dstv1, nil)

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

func DstvVerify(biller, provider string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(dstv2)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, err
	}

	// Create query parameters
	params := url.Values{}
	params.Add("billersCode", biller)
	params.Add("serviceID", provider)

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

func DstvPay(biller, provider, amount, phone, subscription_type, variation_code, request_id string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(dstv3)
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
	params.Add("subscription_type", subscription_type)
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

func Gotv() ([]byte, error) {
	req, _ := http.NewRequest("GET", gotv1, nil)

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

func GotvVerify(biller, provider string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(gotv2)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, err
	}

	// Create query parameters
	params := url.Values{}
	params.Add("billersCode", biller)
	params.Add("serviceID", provider)

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

func GotvPay(biller, provider, amount, phone, subscription_type, request_id string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(gotv3)
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
	params.Add("subscription_type", subscription_type)
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
