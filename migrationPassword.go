package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Password(dbsorint, dbsircles *sql.DB){
	fmt.Println("MIGRATION OF TABLE PASSWORD")

	rows, err := dbsorint.Query("SELECT memberid, password FROM password")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var memberid []byte
	var password []byte

	for rows.Next() {
		err = rows.Scan(&memberid, &password)
		if err != nil {
				fmt.Println(err)
			}
				sqlStatement := `
					INSERT INTO password (memberid, password)
					VALUES ($1, $2)
					RETURNING memberid`

					err = dbsircles.QueryRow(sqlStatement, memberid, password).Scan(&memberid)
	}

	fmt.Println("MIGRATION OF TABLE PASSWORD DONE")
}