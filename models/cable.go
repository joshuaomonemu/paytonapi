package models

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	dstv1    = "https://vtpass.com/api/service-variations?serviceID=dstv"
	dstv2    = "https://vtpass.com/api/merchant-verify"
	dstv3    = "https://vtpass.com/api/pay"
	gotv1    = "https://vtpass.com/api/service-variations?serviceID=gotv"
	gotv2    = "https://vtpass.com/api/merchant-verify"
	gotv3    = "https://vtpass.com/api/pay"
	star1    = "https://vtpass.com/api/service-variations?serviceID=startimes"
	star2    = "https://vtpass.com/api/merchant-verify"
	star3    = "https://vtpass.com/api/pay"
	api      = "e7070f0974c15a7aa6fe5fc6519c5c14"
	public   = "PK_295be67ec6a18b646164ccc9f653adb18c29b5059b0"
	secret   = "SK_4410b58c82cb74f7a2ce604268583f7988041046c8c"
	username = "paytonjit@gmail.com"
	password = "Judithisiayei6561"
)

func Auther() string {
	auth := username + ":" + password
	// Base64 encode the auth string
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	return encodedAuth
}

func Dstv() ([]byte, error) {

	req, _ := http.NewRequest("GET", dstv1, nil)

	req.Header.Add("Authorization", "Basic "+Auther())
	// req.Header.Add("Password", "Judithisiayei6561")

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

	req.Header.Add("Authorization", "Basic "+Auther())
	// req.Header.Add("api-key", api)
	// req.Header.Add("public-key", public)
	// req.Header.Add("secret-key", secret)

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

	// req.Header.Add("api-key", api)
	// req.Header.Add("public-key", public)
	// req.Header.Add("secret-key", secret)
	req.Header.Add("Authorization", "Basic "+Auther())

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

	req.Header.Add("Authorization", "Basic "+Auther())

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

	req.Header.Add("Authorization", "Basic "+Auther())

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

	req.Header.Add("Authorization", "Basic "+Auther())

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

func Star() ([]byte, error) {
	req, _ := http.NewRequest("GET", star1, nil)

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

func StarVerify(biller, provider string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(star2)
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

func StarPay(biller, provider, amount, phone, subscription_type, request_id string) ([]byte, error) {

	// Create a URL object from the base URL
	u, err := url.Parse(star3)
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
