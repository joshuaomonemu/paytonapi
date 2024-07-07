package models

import (
	"app/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// Interswitch base URL
	baseURL = "https://qa.interswitchng.com/quicktellerservice/api/v5/services/categories"
	url1    = "https://qa.interswitchng.com/quicktellerservice/api/v5/services"
	url2    = "https://qa.interswitchng.com/quicktellerservice/api/v5/services/options"
	url3    = "https://qa.interswitchng.com/quicktellerservice/api/v5/Transactions"
	urltt   = "https://sandbox.interswitchng.com/api/v2/quickteller/payments/advices"
	url4    = "https://qa.interswitchng.com/quicktellerservice/api/v5/Transactions/validatecustomers"

	accessToken, _ = auth.GetToken()
)

// Replace with your actual values
var (
	terminalID       = "3pbl0001"
	requestReference = "YFRED3332S3"
	authorization    = accessToken
	bu               = "https://sandbox.interswitchng.com/api/v2/quickteller/payments/advices" // Replace with production URL for live transactions
	contentType      = "application/json"
)

// TransactionData represents the data sent in the request body
type TransactionData struct {
	TerminalID       string  `json:"terminalID"`
	RequestReference string  `json:"requestReference"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	PaymentCode      string  `json:"paymentCode"` // (Optional, required for some transactions)
	// Add other relevant fields as needed based on your transaction type
}

func GetBillersCategories() ([]byte, error) {
	// Construct the API URL
	url := baseURL

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Set the Content-Type header (optional)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("terminalId", "3pbl0001")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check for successful response (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return nil, err
	}

	return body, nil
}

func GetBillersCategoryId(id string) ([]byte, error) {
	// Construct the API URL
	url := url1 + "?categoryId=" + id

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Set the Content-Type header (optional)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("terminalId", "3pbl0001")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check for successful response (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return nil, err
	}

	return body, nil
}

func GetBillerItem(id string) ([]byte, error) {
	// Construct the API URL
	url := url2 + "?serviceid=" + id

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Set the Content-Type header (optional)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("terminalId", "3pbl0001")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check for successful response (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return nil, err
	}

	return body, nil
}

func Advice(item []byte) ([]byte, error) {
	// Construct the API URL
	url := urltt

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(item))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Set the Content-Type header (optional)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("terminalId", "3pbl0001")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check for successful response (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return nil, err
	}

	return body, nil
}

func Validate(item []byte) ([]byte, error) {
	// Construct the API URL
	url := url4

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(item))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Set the Authorization header with the access token
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Set the Content-Type header (optional)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("terminalId", "3pbl0001")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check for successful response (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return nil, err
	}

	return body, nil
}

func Bill() {
	// Prepare transaction data
	transactionData := TransactionData{
		TerminalID:       terminalID,
		RequestReference: requestReference,
		Amount:           100.50, // Replace with your actual amount
		Currency:         "NGN",
	}

	// Marshal data to JSON
	jsonData, err := json.Marshal(transactionData)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, bu, bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", contentType)

	// Send request and handle response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Handle response based on status code and content
	fmt.Println("Response Status:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))
}
