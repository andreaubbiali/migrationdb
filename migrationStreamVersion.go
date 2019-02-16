package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func StreamVersion(dbsorint, dbsircles *sql.DB){
	var maxVersion int
	fmt.Println("MIGRATION OF TABLE STREAMVERSION")

	rows, err := dbsircles.Query("SELECT streamid, category, version FROM event")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var streamid string
	var category string
	var version int
	

	for rows.Next() {
		err = rows.Scan(&streamid, &category, &version)
		if err != nil {
				fmt.Println(err)
			}

			if category == "rolestree"{
				if version > maxVersion{
					maxVersion = version
				}
			}else{
				sqlStatement := `
				INSERT INTO streamversion (streamid, category, version)
				VALUES ($1, $2, $3)
				RETURNING streamid`

				err = dbsircles.QueryRow(sqlStatement, streamid, category, version).Scan(&streamid)
			}
	}
	streamid = "744953eb-ec9f-5f29-9e01-d4ffdd302947"
	category = "rolestree"
	version = maxVersion

	sqlStatement := `
				INSERT INTO streamversion (streamid, category, version)
				VALUES ($1, $2, $3)
				RETURNING streamid`

				err = dbsircles.QueryRow(sqlStatement, streamid, category, version).Scan(&streamid)

	fmt.Println("MIGRATION OF TABLE STREAMVERSION DONE")
}