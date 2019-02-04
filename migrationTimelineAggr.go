//the migration of timeline without the groupID

package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_"github.com/sorintlab/sircles/util"
	_"encoding/hex"
	_"github.com/sorintlab/sircles/common"
	"github.com/satori/go.uuid"
)

//prendere data della tabella event, unmarshallizzare.
//dove il commandtype è qualcosa che crea la timeline importalo come rolestree.
//i member li importi come memberchange(?)
//bisogna importare dagli event tutto anche timeline e commands. 

//PROVIAMO SENZA RIPORTARE LA TIMELINE


//aggregatetype = 'commands' AND eventtype ='CommandExecutionFinished'
// i data sono vuoti

func Timeline(dbsorint, dbsircles *sql.DB){
	var s string
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
	var aggregateid []byte
	var eventtype string
	var data []byte

	

	for rows.Next() {
		err = rows.Scan(&timestamp, &aggregatetype, &eventtype, &aggregateid, &data)
		if err != nil {
				fmt.Println(err)
			}

			if aggregatetype == "timeline"{
				aggregatetype = "rolestree"
			}

			if eventtype == "CommandExecutionFinished"{
				aggregatetype = "rolestree"
			}
				groupid := uuid.Must(uuid.NewV4())
				if aggregatetype == "commands" && eventtype == "CommandExecuted"{
					stringa := fmt.Sprintf(string(data))
					stringaRidotta := stringa[27:60]
					for _ ,x := range stringaRidotta{
						if x != 34{
							s = s + string(x)
						}else{
							break
						}
					}
					switch s{
								case "CreateMember", "UpdateMember", "SetMemberPassword":
									aggregatetype = "memberchange"
								case "CreateTension", "CloseTension":
									aggregatetype = "tension"
								default:
									aggregatetype = "rolestree"
							}
					s = ""
				}
	
				if aggregatetype == "member" && (eventtype == "MemberAvatarSet" || eventtype == "MemberPasswordSet"){
					aggregatetype = "memberchange"
				}
					sqlStatement := `
						INSERT INTO timeline (timestamp, groupid, aggregatetype, aggregateid)
						VALUES ($1, $2, $3, $4)
						RETURNING groupid`
						
					err = dbsircles.QueryRow(sqlStatement, timestamp, groupid, aggregatetype, string(aggregateid)).Scan(&groupid)

			
	}
	fmt.Println("MIGRATION OF TABLE TIMELINE AND AGGREGATEVERSION DONE")
}