package utils

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"strings"
)

func GetOauthToken() *oauth2.Token {
	client := http.Client{}
	credentials := fmt.Sprintf("username=%s&password=%s&grant_type=password", os.Getenv("APIGEE_USERNAME"), os.Getenv("APIGEE_PASSWORD"))
	req, err := http.NewRequest("POST", "https://login.apigee.com/oauth/token", strings.NewReader(credentials))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Hard-coded authorization header by Apigee, refer to https://docs.apigee.com/api-platform/system-administration/management-api-tokens#note
	req.Header.Add("Authorization", "Basic ZWRnZWNsaTplZGdlY2xpc2VjcmV0")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	token := oauth2.Token{}
	_ = json.NewDecoder(res.Body).Decode(&token)
	return &token
}
