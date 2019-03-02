package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func post(token string, body io.Reader) {
	const URLSircles = "http://local.com/upload"
	const ContentType = "application/json"

	// create request
	req, err := http.NewRequest("POST", URLSircles, body)

	// ----- HEADER -----
	// add contenttype
	req.Header.Add("Content-Type", "application/json")
	// add bearer to token
	auth := fmt.Sprintf("Bearer %s", token)
	// add token
	req.Header.Add("authorization", auth)

	// ----- REQUEST -----
	// create client
	client := &http.Client{}
	// do request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("POST error")
		log.Println(err)
	}
	log.Println(resp)
}

func main() {
	// strings.NewReader(s)

}
