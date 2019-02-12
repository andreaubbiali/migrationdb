package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func Member(dbsorint, dbsircles *sql.DB, timeline map[int64]int64) {
	fmt.Println("MIGRATION OF TABLE MEMBER")

	rows, err := dbsorint.Query("SELECT id, start_tl, end_tl, isadmin, username, fullname, email FROM member")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var id []byte
	var start_tl int64
	var end_tl sql.NullInt64
	var isadmin bool
	var username string
	var fullname string
	var email string

	query := `INSERT INTO member (id, start_tl, end_tl, isadmin, username, fullname, email) VALUES `

	values := []interface{}{}
	numFields := 7 // the number of fields you are inserting
	rowsCounts := 0

	for rows.Next() {
		err = rows.Scan(&id, &start_tl, &end_tl, &isadmin, &username, &fullname, &email)
		if err != nil {
			fmt.Println(err)
		}

		for sequence, time := range timeline {
			if sequence == start_tl {
				start_tl = time
			}
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

		if end_tl.Valid {

			for sequence, time := range timeline {
				if sequence == end_tl.Int64 {
					end_tl.Int64 = time
				}
			}

			values = append(values, id, start_tl, end_tl.Int64, isadmin, username, fullname, email)

		} else {
			// append values to query
			values = append(values, id, start_tl, nil, isadmin, username, fullname, email)
		}
	}
	// remove last ','
	query = query[:len(query)-1]
	// execute query
	_, err = dbsircles.Exec(query, values...)
	if err != nil {
		log.Println("Query error")
		log.Println(err)
	}
	fmt.Println("MIGRATION OF TABLE MEMBER DONE")
}
