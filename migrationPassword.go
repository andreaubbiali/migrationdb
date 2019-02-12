package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func Password(dbsorint, dbsircles *sql.DB) {
	fmt.Println("MIGRATION OF TABLE PASSWORD")

	rows, err := dbsorint.Query("SELECT memberid, password FROM password")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var memberid []byte
	var password []byte

	query := `INSERT INTO password (memberid, password)	VALUES `

	values := []interface{}{}
	numFields := 2 // the number of fields you are inserting
	rowsCounts := 0

	for rows.Next() {
		err = rows.Scan(&memberid, &password)
		if err != nil {
			fmt.Println(err)
		}

		// count value to insert
		n := rowsCounts * numFields
		rowsCounts++

		// values insert
		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		// remove last ','
		query = query[:len(query)-1] + `),`

		// append values to query
		values = append(values, memberid, password)
	}
	// remove last ','
	query = query[:len(query)-1]
	// execute query
	_, err = dbsircles.Exec(query, values...)
	if err != nil {
		log.Println("Query error")
		log.Println(err)
	}
	fmt.Println("MIGRATION OF TABLE PASSWORD DONE")
}
