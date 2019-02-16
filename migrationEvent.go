package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type metadata struct {
	Correlationid, CausationID, GroupID, CommandIssuerID string
}

func Event(dbsorint, dbsircles *sql.DB) {
	var meta metadata
	var s string
	var n int

	fmt.Println("MIGRATION OF TABLE EVENT")

	rows, err := dbsorint.Query("SELECT id, sequencenumber, eventtype, aggregatetype, aggregateid, timestamp, version, correlationid, causationid, data FROM event ORDER BY sequencenumber")
	//rows, err := dbsorint.Query("SELECT id, sequencenumber, eventtype, aggregatetype, aggregateid, timestamp, version, correlationid, causationid, data FROM event WHERE aggregatetype != 'commands' AND aggregatetype != 'timeline'")
	if err != nil {
		fmt.Println("a", err)
	}
	defer rows.Close()

	var id []byte
	var sequencenumber int
	var eventtype string
	var aggregatetype string
	var aggregateid []byte
	var timestamp string
	var version int
	var data []byte

	query := `INSERT INTO event (id, eventtype, category, streamid, timestamp, version, data, metadata) values `

	values := []interface{}{}
	numFields := 8 // the number of fields you are inserting
	rowsCounts := 0

	for rows.Next() {
		err = rows.Scan(&id, &sequencenumber, &eventtype, &aggregatetype, &aggregateid, &timestamp, &version, &meta.Correlationid, &meta.CausationID, &data)
		if err != nil {
			fmt.Println(err)
		}
		// stringa := string(timestamp)
		// time := fmt.Sprintf(stringa[:10] + " " + stringa[11:29])
		str := fmt.Sprintf(`SELECT groupid FROM timeline WHERE timestamp = '%v'`, timestamp)
		sqlStatement := str
		row := dbsircles.QueryRow(sqlStatement)

		err := row.Scan(&meta.GroupID)
		if err != nil {
			fmt.Println(timestamp)
			fmt.Println(err)
			//fmt.Println(aggregatetype, eventtype)
		}

		if aggregatetype == "timeline" {
			aggregatetype = "rolestree"
			eventtype = "RoleUpdated"
		}

		if eventtype == "CommandExecutionFinished" {
			aggregatetype = "rolestree"
			eventtype = "RoleUpdated"
		}
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
				eventtype = "MemberUpdated"
			case "CreateTension", "CloseTension":
				aggregatetype = "tension"
				eventtype = "TensionCreated"
			default:
				aggregatetype = "rolestree"
				eventtype = "RoleUpdated"
			}
			s = ""
		}

		groupMetadata := metadata{
			Correlationid:   meta.Correlationid,
			CausationID:     meta.CausationID,
			GroupID:         meta.GroupID,
			CommandIssuerID: "00000000-0000-0000-0000-000000000000",
		}

		marshalmetadata, _ := json.Marshal(groupMetadata)

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
		values = append(values, id, eventtype, aggregatetype, aggregateid, timestamp, version, data, marshalmetadata)
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

			query = `INSERT INTO event (id, eventtype, category, streamid, timestamp, version, data, metadata) values `
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
	fmt.Println("MIGRATION OF TABLE EVENT DONE")
}
