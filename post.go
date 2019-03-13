package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Post do a post request to sircles api
func Post(token string, body io.Reader) {
	const ContentType = "application/json"
	const URLSircles = "http://localhost:8080/api/graphql"

	// create post request
	req, err := http.NewRequest("POST", URLSircles, body)
	if err != nil {
		log.Println("Error create post request")
		log.Println(err)
	}

	// header
	// add contenttype
	req.Header.Add("Content-Type", "application/json")
	// add bearer to token
	auth := fmt.Sprintf("Bearer %s", token)
	// add token
	req.Header.Add("authorization", auth)

	// request
	// create client
	client := &http.Client{}
	// do request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("POST error")
		log.Println(err)
		respParsed, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(respParsed))
	}
}
