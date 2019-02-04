package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Member(dbsorint, dbsircles *sql.DB, timeline map[int64]int64){
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

	for rows.Next() {
		err = rows.Scan(&id, &start_tl, &end_tl, &isadmin, &username, &fullname, &email)
		if err != nil {
				fmt.Println(err)
			}
	
			for sequence, time := range timeline{
				if sequence == start_tl{
					start_tl = time
				}
			}
	
			if end_tl.Valid {
				
				for sequence, time := range timeline{
					if sequence == end_tl.Int64{
						end_tl.Int64 = time
					}
				}
				sqlStatement := `
					INSERT INTO member (id, start_tl, end_tl, isadmin, username, fullname, email)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, end_tl.Int64, isadmin, username, fullname, email).Scan(&id)

			} else{
					sqlStatement := `
					INSERT INTO member (id, start_tl, isadmin, username, fullname, email)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, isadmin, username, fullname, email).Scan(&id)
			}
	}
	fmt.Println("MIGRATION OF TABLE MEMBER DONE")
}
