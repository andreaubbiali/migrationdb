//the migration of timeline without the groupID

package main

import (
	"database/sql"
	_ "encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	_ "github.com/sorintlab/sircles/common"
	_ "github.com/sorintlab/sircles/util"
)

// migration events from roleevent in timeline
// circlechangesapplies events
func RoleEventTimeline(dbsorint, dbsircles *sql.DB) {
	var n int

	fmt.Println("MIGRATION OF TABLE ROLE EVENT AND TIMELINE")

	rows, err := dbsorint.Query("SELECT t.timestamp FROM roleevent r, timeline t WHERE t.sequencenumber = r.timeline")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var timestamp time.Time
	var aggregatetype string
	var aggregateid string

	query := `INSERT INTO timeline (timestamp, groupid, aggregatetype, aggregateid) VALUES `

	values := []interface{}{}
	numFields := 4 // the number of fields you are inserting
	rowsCounts := 0

	for rows.Next() {
		err = rows.Scan(&timestamp)
		if err != nil {
			fmt.Println(err)
		}
		// generate group id
		groupid := uuid.Must(uuid.NewV4())
		// define aggregate
		aggregatetype = "rolestree"
		aggregateid = "744953eb-ec9f-5f29-9e01-d4ffdd302947"

		// count value to insert
		if rowsCounts == 0 {
			n = 0
		} else {
			n = rowsCounts * numFields
		}

		rowsCounts++

		// values insert
		query += `(`
		for j := 0; j < numFields; j++ {
			query += `$` + strconv.Itoa(n+j+1) + `,`
		}
		// remove last ','
		query = query[:len(query)-1] + `),`

		// append values to query
		values = append(values, timestamp, groupid, aggregatetype, aggregateid)

		// max number of values in postgresql is 65535.
		if len(values) >= 65535-numFields {
			// remove last ','
			query = query[:len(query)-1]
			// execute query
			_, err = dbsircles.Exec(query, values...)
			if err != nil {
				log.Println("Query error")
				log.Println(err)
			}
			// prepare new insert
			query = `INSERT INTO timeline (timestamp, groupid, aggregatetype, aggregateid) VALUES `
			rowsCounts = 0
			values = nil
		}

	}
	// check if there are others rows to insert
	if rowsCounts > 0 {
		// remove last ','
		query = query[:len(query)-1]
		// execute query
		_, err = dbsircles.Exec(query, values...)
		if err != nil {
			log.Println("Query error")
			log.Println(err)
		}
	}
	fmt.Println("MIGRATION OF TABLE ROLE EVENT AND TIMELINE DONE")
}
