//the migration of timeline without the groupID

package main

import (
	"database/sql"
	_ "encoding/hex"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	_ "github.com/sorintlab/sircles/common"
	_ "github.com/sorintlab/sircles/util"
)

//prendere data della tabella event, unmarshallizzare.
//dove il commandtype è qualcosa che crea la timeline importalo come rolestree.
//i member li importi come memberchange(?)
//bisogna importare dagli event tutto anche timeline e commands.

//PROVIAMO SENZA RIPORTARE LA TIMELINE

//aggregatetype = 'commands' AND eventtype ='CommandExecutionFinished'
// i data sono vuoti

func Timeline(dbsorint, dbsircles *sql.DB) {
	var s string
	var n int

	fmt.Println("MIGRATION OF TABLE TIMELINE AND AGGREGATEVERSION")

	rows, err := dbsorint.Query("SELECT timestamp, aggregatetype, eventtype, aggregateid, data  FROM event")
	//rows, err := dbsorint.Query("SELECT timestamp, aggregatetype, eventtype, aggregateid, data  FROM event WHERE aggregatetype != 'timeline'")
	//	rows, err := dbsorint.Query("SELECT timestamp, aggregatetype, eventtype, aggregateid  FROM event WHERE aggregatetype = 'member' OR aggregatetype = 'rolestree' OR aggregatetype = 'tension'")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var timestamp string
	var aggregatetype string
	var aggregateid string
	var eventtype string
	var data []byte

	query := `INSERT INTO timeline (timestamp, groupid, aggregatetype, aggregateid) VALUES `

	values := []interface{}{}
	numFields := 4 // the number of fields you are inserting
	rowsCounts := 0

	for rows.Next() {
		err = rows.Scan(&timestamp, &aggregatetype, &eventtype, &aggregateid, &data)
		if err != nil {
			fmt.Println(err)
		}

		if aggregatetype == "timeline" {
			aggregatetype = "rolestree"
		}

		if eventtype == "CommandExecutionFinished" {
			aggregatetype = "rolestree"
		}
		groupid := uuid.Must(uuid.NewV4())
		if aggregatetype == "commands" && eventtype == "CommandExecuted" {
			stringa := fmt.Sprintf(string(data))
			stringaRidotta := stringa[27:60]
			for _, x := range stringaRidotta {
				if x != 34 {
					s = s + string(x)
				} else {
					break
				}
			}
			switch s {
			case "CreateMember", "UpdateMember", "SetMemberPassword":
				aggregatetype = "memberchange"
			case "CreateTension", "CloseTension":
				aggregatetype = "tension"
			default:
				aggregatetype = "rolestree"
			}
			s = ""
		}

		if aggregatetype == "member" && (eventtype == "MemberAvatarSet" || eventtype == "MemberPasswordSet") {
			aggregatetype = "memberchange"
		}

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
	fmt.Println("MIGRATION OF TABLE TIMELINE AND AGGREGATEVERSION DONE")
}
