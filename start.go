package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// BodyRequest for sircles api
type BodyRequest struct {
	OperationName string                 `json:"operationName"`
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

func main() {
	// strings.NewReader(s)
	// create connection to db sorint
	dbsorint, err := sql.Open("postgres", "postgres://postgres:password@localhost/sorint?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer dbsorint.Close()

	token := Auth()

	// member
	Member(dbsorint, token)
}
