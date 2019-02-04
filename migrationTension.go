package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Tension(dbsorint, dbsircles *sql.DB, timeline map[int64]int64){
	fmt.Println("MIGRATION OF TABLE TENSION")

	rows, err := dbsorint.Query("SELECT id, start_tl, end_tl, title, description, closed, closereason FROM tension")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var id []byte
	var start_tl int64
	var end_tl sql.NullInt64
	var title string
	var description string
	var closed bool
	var closereason sql.NullString
	

	for rows.Next() {
		err = rows.Scan(&id, &start_tl, &end_tl, &title, &description, &closed, &closereason)
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

				if closereason.Valid{
						sqlStatement := `
					INSERT INTO tension (id, start_tl, end_tl, title, description, closed, closereason)
					VALUES ($1, $2, $3, $4, $5, $6, $7)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, end_tl.Int64, title, description, closed, closereason.String).Scan(&id)
				}else{
						sqlStatement := `
					INSERT INTO tension (id, start_tl, end_tl, title, description, closed)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, end_tl.Int64, title, description, closed, closereason).Scan(&id)
				}
			

			} else{
				if closereason.Valid{
						sqlStatement := `
					INSERT INTO tension (id, start_tl, title, description, closed, closereason)
					VALUES ($1, $2, $3, $4, $5, $6)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, title, description, closed, closereason.String).Scan(&id)
				}else{
						sqlStatement := `
					INSERT INTO tension (id, start_tl, title, description, closed)
					VALUES ($1, $2, $3, $4, $5)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, title, description, closed).Scan(&id)
				}	
			}
	}
	fmt.Println("MIGRATION OF TABLE TENSION DONE")
}