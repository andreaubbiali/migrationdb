package main

import (
	"database/sql"
	_ "encoding/hex"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/satori/go.uuid"
	_ "github.com/sorintlab/sircles/common"
	_ "github.com/sorintlab/sircles/util"
)

func Try(dbsorint, dbsircles *sql.DB) {
	var s string

	rows, err := dbsorint.Query("SELECT timestamp, aggregatetype, eventtype, aggregateid, data  FROM event WHERE aggregatetype != 'timeline'")
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
				aggregatetype = "member"
			case "CreateTension", "CloseTension":
				aggregatetype = "tension"
			default:
				aggregatetype = "rolestree"
			}
			s = ""
		}

	}

}

// //rolestree
// RoleAddMember
// CircleUnsetLeadLinkMember
// UpdateRootRole
// CircleUpdateChildRole
// SetupRootRole
// CircleCreateChildRole
// CircleSetLeadLinkMember
// CircleSetCoreRoleMember
// CircleUnsetCoreRoleMember
// SetRoleAdditionalContent
// RoleRemoveMember
// CircleDeleteChildRole

// //member
// CreateMember
// UpdateMember
// SetMemberPassword

// //tension
// CreateTension
// CloseTension
