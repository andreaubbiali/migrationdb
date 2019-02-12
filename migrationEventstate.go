//bisogna prendere il numero della tabella eventstate e fare un for che butti dentro da 1 a quel numero nella tabellea
//sequencenumber di sircles. operaione moooolto lunga
package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func EventState(dbsorint, dbsircles *sql.DB) {
	fmt.Println("MIGRATION OF TABLE EVENTSTATE")

	sqlStatement := `SELECT * FROM eventstate;`
	var number int64
	var j int64

	query := `INSERT INTO sequencenumber (sequencenumber) VALUES `

	values := []interface{}{}
	numFields := 1 // the number of fields you are inserting
	rowsCounts := 0

	row := dbsorint.QueryRow(sqlStatement)
	err := row.Scan(&number)
	if err != nil {
		fmt.Println(err)
	}

	for j = 1; j <= number; j++ {

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
		values = append(values, j)
	}
	// remove last ','
	query = query[:len(query)-1]
	// execute query
	_, err = dbsircles.Exec(query, values...)
	if err != nil {
		log.Println("Query error")
		log.Println(err)
	}
	fmt.Println("MIGRATION OF TABLE EVENTSTATE DONE")
}
