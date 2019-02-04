package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func RoleAdditionalContent(dbsorint, dbsircles *sql.DB, timeline map[int64]int64){
	fmt.Println("MIGRATION OF TABLE ROLEADDITIONALCONTENT")

	rows, err := dbsorint.Query("SELECT id, start_tl, end_tl, content FROM roleadditionalcontent")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var id []byte
	var start_tl int64
	var end_tl sql.NullInt64
	var content string
	

	for rows.Next() {
		err = rows.Scan(&id, &start_tl, &end_tl, &content)
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
					INSERT INTO roleadditionalcontent (id, start_tl, end_tl, content)
					VALUES ($1, $2, $3, $4)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, end_tl.Int64,content).Scan(&id)

			} else{
					sqlStatement := `
					INSERT INTO roleadditionalcontent (id, start_tl, content)
					VALUES ($1, $2, $3)
					RETURNING id`

					err = dbsircles.QueryRow(sqlStatement, id, start_tl, content).Scan(&id)
			}
	}
	fmt.Println("MIGRATION OF TABLE ROLEADDITIONALCONTENT DONE")
}