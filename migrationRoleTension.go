package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func RoleTension(dbsorint, dbsircles *sql.DB, timeline map[int64]int64){
	fmt.Println("MIGRATION OF TABLE ROLETENSION")

	rows, err := dbsorint.Query("SELECT start_tl, end_tl, x, y FROM roletension")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var start_tl int64
	var end_tl sql.NullInt64
	var x []byte
	var y []byte
	

	for rows.Next() {
		err = rows.Scan(&start_tl, &end_tl, &x, &y)
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
					INSERT INTO roletension (start_tl, end_tl, x, y)
					VALUES ($1, $2, $3, $4)
					RETURNING x`

					err = dbsircles.QueryRow(sqlStatement, start_tl, end_tl.Int64, x, y).Scan(&x)

			} else{
					sqlStatement := `
					INSERT INTO roletension (start_tl, x, y)
					VALUES ($1, $2, $3)
					RETURNING x`

					err = dbsircles.QueryRow(sqlStatement, start_tl, x, y).Scan(&x)
			}
	}
	fmt.Println("MIGRATION OF TABLE ROLETENSION DONE")
}