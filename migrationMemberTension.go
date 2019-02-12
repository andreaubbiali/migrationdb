package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func MemberTension(dbsorint, dbsircles *sql.DB, timeline map[int64]int64) {
	fmt.Println("MIGRATION OF TABLE MEMBERTENSION")

	rows, err := dbsorint.Query("SELECT start_tl, end_tl, x, y FROM membertension")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var start_tl int64
	var end_tl sql.NullInt64
	var x []byte
	var y []byte

	query := `INSERT INTO membertension (start_tl, end_tl, x, y) VALUES `

	values := []interface{}{}
	numFields := 4 // the number of fields you are inserting
	rowsCounts := 0

	for rows.Next() {
		err = rows.Scan(&start_tl, &end_tl, &x, &y)
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

			// append values to query
			values = append(values, start_tl, end_tl.Int64, x, y)

		} else {
			// append values to query
			values = append(values, start_tl, nil, x, y)
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
	fmt.Println("MIGRATION OF TABLE MEMBERTENSION DONE")
}
