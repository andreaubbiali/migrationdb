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

	/* ----- CREATE CONNECTION TO DB ----- */

	fmt.Println("CONNECTION TO DB")

	// connection to database sorint
	dbsorint, err := sql.Open("postgres", "postgres://postgres:password@localhost/sorint?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer dbsorint.Close()

	// connection to database sircles
	dbsircles, err := sql.Open("postgres", "postgres://sircles:password@localhost/sircles?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer dbsircles.Close()

	fmt.Println("CONNECTION TO DB DONE")

	/* ----- DROP AND CREATE TABLE ----- */

	/*
		fmt.Println("CREATE TABLE")

		// read query
		file, err := ioutil.ReadFile("./tables.sql")
		if err != nil {
			fmt.Println("Errors create table")
		}

		// split query divided by ';'
		requests := strings.Split(string(file), ";")

		// execute each query saved in the file
		for _, request := range requests {
			dbsircles.Exec(request)
		}

		fmt.Println("CREATE TABLE DONE")


	/* ----- SEND DATA TO API ----- */

	fmt.Println("API REQUESTS")

	token := Auth()

	// not used, members are imported from Clann db

	fmt.Println("MEMBER API")
	//Member(dbsorint, token)
	fmt.Println("MEMBER API DONE")

	fmt.Println("ROLE API")
	Role(dbsorint, token)
	fmt.Println("ROLE API DONE")
	// RoleMember(dbsorint, token)
	// Domain(dbsorint, token)
	// Accountabilities(dbsorint, token)

	fmt.Println("API REQUESTS DONE")
}
