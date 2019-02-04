//bisogna prendere il numero della tabella eventstate e fare un for che butti dentro da 1 a quel numero nella tabellea 
//sequencenumber di sircles. operaione moooolto lunga
package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func EventState(dbsorint, dbsircles *sql.DB){
	fmt.Println("MIGRATION OF TABLE EVENTSTATE")

	sqlStatement := `SELECT * FROM eventstate;`
	var number int64
	var j int64

	row := dbsorint.QueryRow(sqlStatement)
	err := row.Scan(&number)
	if err != nil{
		fmt.Println(err)
	}
	
	for j = 1 ; j <= number ; j++{
		sqlStatement := `
					INSERT INTO sequencenumber (sequencenumber)
					VALUES ($1)
					RETURNING sequencenumber`
	
					err = dbsircles.QueryRow(sqlStatement, j).Scan(&j)
	}
	fmt.Println("MIGRATION OF TABLE EVENTSTATE DONE")
}