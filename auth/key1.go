package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Access struct {
	Access_Token string `json:"access_token"`
}

func Auth2() (string, error) {

	url := "https://auth.reloadly.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"LrSJMJirX0iR3MB6AKj9G2P6Ui3YsXjO\",\"client_secret\":\"ydGmgekGSB-8bj8mQuB5ElAdMMWm3A-e175DalsnwHRqwLb5R51Yf6i65dXwLUV\",\"grant_type\":\"client_credentials\",\"audience\":\"https://utilities-sandbox.reloadly.com\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var access *Access

	err = json.Unmarshal(body, &access)
	if err != nil {
		return "", err
	}
	return access.Access_Token, nil

}
