package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// Replace with your Interswitch Client ID
	clientID = "IKIA72C65D005F93F30E573EFEAC04FA6DD9E4D344B1"
	// Replace with your Interswitch Secret Key
	secretKey = "YZMqZezsltpSPNb4+49PGeP7lYkzKn1a5SaVSyzKOiI="
	// Interswitch OAuth endpoint (test environment)
	oauthURL = "https://passport.k8.isw.la/passport/oauth/token"
)

func GetToken() (string, error) {
	// Combine Client ID and Secret Key separated by colon
	credentials := []byte(clientID + ":" + secretKey)

	// Base64 encode the credentials
	encodedCredentials := base64.StdEncoding.EncodeToString(credentials)

	// Create the HTTP request
	req, err := http.NewRequest(http.MethodPost, oauthURL, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Set the Authorization header with Basic authentication and encoded credentials
	req.Header.Set("Authorization", "Basic "+encodedCredentials)

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Set the request body with grant type parameter
	requestBody := []byte("grant_type=client_credentials")
	req.Body = ioutil.NopCloser(bytes.NewReader(requestBody))

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Check for successful response (status code 200)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode, string(body))
		return "", err
	}

	// Unmarshal the JSON response
	var tokenResponse map[string]interface{}
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	// Extract the access token
	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		err1 := fmt.Errorf("access_token not found in response")
		return "", err1
	}

	return accessToken, nil
}
