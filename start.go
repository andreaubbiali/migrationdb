///////////////////////////////////////////////////// public | accountability        | table | sircles
///////////////////////////////////////////////////// public | aggregateversion      | table | sircles
///////////////////////////////////////////////////// public | circledirectmember    | table | sircles//è vuota////////////////////////
////////////////////////////////////////////////////// public | commandevent          | table | sircles//è vuota ////////////////////
////////////////////////////////////////////////////// public | domain                | table | sircles
///////////////////////////////////////////////////// public | event                 | table | sircles
////////////////////////////////////////////////////// public | eventstate            | table | sircles
///////////////////////////////////////////////////// public | member                | table | sircles
//////////////////////////////////////////////////// public | memberavatar          | table | sircles
//////////////////////////////////////////////////// public | memberevent           | table | sircles//è vuota/////////////////
//////////////////////////////////////////////////// public | membermatch           | table | sircles// è vuota /////////
//////////////////////////////////////////////////// public | membertension         | table | sircles
// public | migration             | table | sircles
/////////////////////////////////////////////////// public | password              | table | sircles
//////////////////////////////////////////////////// public | role                  | table | sircles
//////////////////////////////////////////////////// public | roleaccountability    | table | sircles
/////////////////////////////////////////////////// public | roleadditionalcontent | table | sircles
////////////////////////////////////////////////// public | roledomain            | table | sircles
/////////////////////////////////////////////////// public | roleevent             | table | sircles
/////////////////////////////////////////////////// public | rolemember            | table | sircles
/////////////////////////////////////////////////// public | rolerole              | table | sircles
/////////////////////////////////////////////////// public | roletension           | table | sircles
////////////////////////////////////////////////// public | tension               | table | sircles
////////////////////////////////////////////// public | timeline              | table | sircles

//12  minuti

//SE EVENTSTATE FOSSE DI UN SOLO NUMERO PIÙ GRANDE O PIÙ PICCOLO POTREBBE CREARE PROBELI. GUARDACI.
//PRIMA DI FAR PARTIRE OGNI COSA ELIMINA I RECORD
//ALLA FINE ELIMINI TUTTO SB, LO RICREI FAI PARTIRE HOST880. E CANCELLI I RECORD E FAI PARTIRE LA MIGRATION. oppure fai le create table?
//sopra le pagine ho scritto se e cosa potrebbe dare problemi


//C'È UN PROBLEMA NEL RUOLO PADRE E FIGLIO. FIN CHE SI CLICCA SUI CHILD VA TUTTO BNE MA QUANDO VUOI ENTRARE NEL RUOLO PADRE MANCANO DEI COLLEGAMENTI.
// aaaasircles -->>>>>>>>>>>>>>>>> nome del database sircles migrato

//popolare la tabella streamversion
package main

import(
	"fmt"
	_ "github.com/lib/pq"
	"time"
	"database/sql"
)


func main(){
	
	var timeline map[int64]int64
	timeline = make(map[int64]int64)

	//connection to database sorint.
	dbsorint, err := sql.Open("postgres", "postgres://sircles:password@localhost/sorint?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer dbsorint.Close()

	//connection to database sircles.
	dbsircles, err := sql.Open("postgres", "postgres://postgres:password@localhost/sircles?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer dbsircles.Close()


	//INFORMATION NEEDED FROM TIMELINE
	rows, err := dbsorint.Query("SELECT sequencenumber, timestamp FROM timeline ")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var sequencenumber int64
	var timestamp time.Time

	for rows.Next() {
		err = rows.Scan(&sequencenumber, &timestamp)
		if err != nil {
			fmt.Println(err)
		}
		timeline[sequencenumber] = timestamp.UnixNano()
	}

	//every time I do a query of INSERT INTO i must use also a scan and returning something because without it the code after some loop broke.

	// //ACCOUNTABILITY
	// Accountability(dbsorint, dbsircles, timeline)

	// //DOMAIN
	// Domain(dbsorint, dbsircles, timeline)

	// //EVENTSTATE
	// EventState(dbsorint, dbsircles)

	// //MEMBER
	// Member(dbsorint, dbsircles, timeline)

	// //MEMBERAVATAR
	// MemberAvatar(dbsorint, dbsircles, timeline)

	// //MEMBERTENSION
	// MemberTension(dbsorint, dbsircles, timeline)

	// //PASSWORD
	// Password(dbsorint, dbsircles)

	// //ROLE
	// Role(dbsorint, dbsircles, timeline)

	// //ROLEACCOUNTABILITY
	// RoleAccountability(dbsorint, dbsircles, timeline)

	// //ROLEADDITIONALCONTENT
	// RoleAdditionalContent(dbsorint, dbsircles, timeline)

	// //ROLEDOMAIN
	// RoleDomain(dbsorint, dbsircles, timeline)

	// //ROLEMEMBER
	// RoleMember(dbsorint, dbsircles, timeline)

	// //ROLEROLE
	// RoleRole(dbsorint, dbsircles, timeline)

	// //ROLETENSION
	// RoleTension(dbsorint, dbsircles, timeline)

	// //TENSION
	// Tension(dbsorint, dbsircles, timeline)

	// //ROLEEVENT
	// RoleEvent(dbsorint, dbsircles, timeline)

	//TIMELINE AND AGGREGATEVERSION 
	Timeline(dbsorint, dbsircles)

	//EVENT
	Event(dbsorint, dbsircles)

	// //TRY
	// Try(dbsorint, dbsircles)

}


//go get -u github.com/lib/pq (per scaricare il pacchetto)

