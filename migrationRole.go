//potrebbe dare problemi per il nome dei roletype

package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Role(dbsorint, dbsircles *sql.DB, timeline map[int64]int64){
	fmt.Println("MIGRATION OF TABLE ROLE")

	rows, err := dbsorint.Query("SELECT id, start_tl, end_tl, roletype, depth, name, purpose FROM role")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var id []byte
	var start_tl int64
	var end_tl sql.NullInt64
	var roletype string
	var depth int
	var name string
	var purpose string

	for rows.Next() {
		err = rows.Scan(&id, &start_tl, &end_tl, &roletype, &depth, &name, &purpose)
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
					INSERT INTO role (id, start_tl, end_tl, roletype, depth, name, purpose)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, end_tl.Int64, roletype, depth, name, purpose).Scan(&id)

			} else{
					sqlStatement := `
					INSERT INTO role (id, start_tl, roletype, depth, name, purpose)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, roletype, depth, name, purpose).Scan(&id)
			}
	}
	fmt.Println("MIGRATION OF TABLE MEMBERTENSION DONE")
}
