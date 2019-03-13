package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Auth api authentication
func Auth() string {
	var body []byte
	var parsedBody map[string]interface{}
	var token string
	// make login url
	const loginURL = "http://localhost:8080/api/auth/login"
	// make post request
	resp, err := http.PostForm(loginURL,
		url.Values{"login": {"admin"}, "password": {"password"}})
	if err != nil {
		log.Println("Error auth request")
		log.Println(err)
	}
	// read body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error auth request")
		log.Println(err)
	}
	// unmarshall body response
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		log.Println("Error unmarshall auth request")
		log.Println(err)
	}
	// read token
	token = parsedBody["token"].(string)

	return token
}
