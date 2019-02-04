package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func RoleEvent(dbsorint, dbsircles *sql.DB, timeline map[int64]int64){
	fmt.Println("MIGRATION OF TABLE ROLEEVENT")

	rows, err := dbsorint.Query("SELECT timeline, id, command, cause, eventtype, roleid, data FROM roleevent")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var timel int64
	var id []byte
	var command sql.NullString
	var cause sql.NullString
	var eventtype string
	var roleid []byte
	var data []byte
	

	for rows.Next() {
		err = rows.Scan(&timel, &id, &command, &cause, &eventtype, &roleid, &data)
		if err != nil {
			fmt.Println(err)
		}
	
		for sequence, time := range timeline{
			if sequence == timel{
				timel = time
			}
		}
				if command.Valid{

					if cause.Valid{
							sqlStatement := `
						INSERT INTO roleevent (timeline, id, command, cause, eventtype, roleid, data)
						VALUES ($1, $2, $3, $4, $5, $6, $7)
						RETURNING id`
	
						err = dbsircles.QueryRow(sqlStatement, timel, id, command, cause, eventtype, roleid, data).Scan(&id)
					}else{
							sqlStatement := `
						INSERT INTO roleevent (timeline, id, command, eventtype, roleid, data)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING id`

						err = dbsircles.QueryRow(sqlStatement, timel, id, command, eventtype, roleid, data).Scan(&id)
					}
						
				}else{
					
					if cause.Valid{
							sqlStatement := `
						INSERT INTO roleevent (timeline, id, cause, eventtype, roleid, data)
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING id`
	
						err = dbsircles.QueryRow(sqlStatement, timel, id, cause, eventtype, roleid, data).Scan(&id)
					}else{
							sqlStatement := `
						INSERT INTO roleevent (timeline, id, eventtype, roleid, data)
						VALUES ($1, $2, $3, $4, $5)
						RETURNING id`
	
						err = dbsircles.QueryRow(sqlStatement, timel, id, eventtype, roleid, data).Scan(&id)
					}
						
				}
			

	}
	fmt.Println("MIGRATION OF TABLE ROLEEVENT DONE")
}