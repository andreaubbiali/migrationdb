package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func RoleMember(dbsorint, dbsircles *sql.DB, timeline map[int64]int64){
	fmt.Println("MIGRATION OF TABLE ROLEMEMBER")

	rows, err := dbsorint.Query("SELECT start_tl, end_tl, x, y, focus, nocoremember, electionexpiration FROM rolemember")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var start_tl int64
	var end_tl sql.NullInt64
	var x []byte
	var y []byte
	var focus sql.NullString
	var nocoremember bool
	var electionexpiration []byte
	

	for rows.Next() {
		err = rows.Scan(&start_tl, &end_tl, &x, &y, &focus, &nocoremember, &electionexpiration)
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

				if focus.Valid{

					if electionexpiration == nil{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, end_tl, x, y, focus, nocoremember)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING x`

						err = dbsircles.QueryRow(sqlStatement, start_tl, end_tl.Int64, x, y, focus.String, nocoremember).Scan(&x)

					}else{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, end_tl, x, y, focus, nocoremember, electionexpiration)
						VALUES ($1, $2, $3, $4, $5, $6, $7)
						RETURNING x`

						err = dbsircles.QueryRow(sqlStatement, start_tl, end_tl.Int64, x, y, focus.String, nocoremember, electionexpiration).Scan(&x)
					}
					

				}else{

					if electionexpiration == nil{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, end_tl, x, y, nocoremember)
						VALUES ($1, $2, $3, $4, $5)
						RETURNING x`

						err = dbsircles.QueryRow(sqlStatement, start_tl, end_tl.Int64, x, y, nocoremember).Scan(&x)
					}else{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, end_tl, x, y, nocoremember, electionexpiration)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING x`

						err = dbsircles.QueryRow(sqlStatement, start_tl, end_tl.Int64, x, y, nocoremember, electionexpiration).Scan(&x)
					}
					
				}
				

					
			} else{
				if focus.Valid{
					if electionexpiration == nil{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, x, y, focus, nocoremember)
						VALUES ($1, $2, $3, $4, $5)
						RETURNING x`

						err = dbsircles.QueryRow(sqlStatement, start_tl, x, y, focus.String, nocoremember).Scan(&x)
					}else{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, x, y, focus, nocoremember, electionexpiration)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING x`

						err = dbsircles.QueryRow(sqlStatement, start_tl, x, y, focus.String, nocoremember, electionexpiration).Scan(&x)
					}
					

				}else{
					if electionexpiration == nil{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, x, y, nocoremember)
						VALUES ($1, $2, $3, $4)
						RETURNING x`

						err = dbsircles.QueryRow(sqlStatement, start_tl, x, y, nocoremember).Scan(&x)
					}else{
							sqlStatement := `
						INSERT INTO rolemember (start_tl, x, y, nocoremember, electionexpiration)
						VALUES ($1, $2, $3, $4, $5)
						RETURNING x`
	
						err = dbsircles.QueryRow(sqlStatement, start_tl, x, y, nocoremember, electionexpiration).Scan(&x)
					}
					
				}
					
			}
	}
	fmt.Println("MIGRATION OF TABLE ROLEMEMBER DONE")
}