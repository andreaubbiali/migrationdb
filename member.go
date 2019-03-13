package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

// CreateMemberChange
type CreateMemberChange struct {
	Email    string `json:"email"`
	Fullname string `json:"fullName"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}

// Member make api request to add members
func Member(dbsorint *sql.DB, token string) {
	var member CreateMemberChange
	var body BodyRequest

	// add body information about api request
	body.OperationName = "createMember"
	body.Query = "mutation createMember($createMemberChange: CreateMemberChange!) {\n  createMember(createMemberChange: $createMemberChange) {\n    hasErrors\n    genericError\n    createMemberChangeErrors {\n      userName\n      fullName\n      email\n      password\n      __typename\n    }\n    member {\n      uid\n      isAdmin\n      userName\n      fullName\n      email\n      __typename\n    }\n    __typename\n  }\n}\n"

	// member query
	rows, err := dbsorint.Query("SELECT username, fullname, email FROM member WHERE end_tl IS NULL")
	if err != nil {
		log.Println("Member query error")
		log.Println(err)
	}

	for rows.Next() {
		// read values
		_ = rows.Scan(&member.UserName, &member.Fullname, &member.Email)
		// set password
		member.Password = fmt.Sprintf("%sprova", member.UserName)
		// initialize map
		body.Variables = make(map[string]interface{})
		// set member body variables
		body.Variables["createMemberChange"] = member
		// make json body
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			log.Println("Error member marshall")
			log.Println(err)
		}
		// make reader body
		bodyReader := bytes.NewReader(bodyJSON)
		Post(token, bodyReader)
	}
}
