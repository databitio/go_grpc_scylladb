package queries

import (
	// "encoding/json"
	"fmt"
	"log"

	// "errors"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)


type ticket struct {
	Ticketid    gocql.UUID `json:"ticketid"`
	Title       string `json:"title"`
	Userid    	string `json:"userid"`
	Description string `json:"description"`
	Reward      string `json:"reward"`
	Lifespan    int `json:"lifespan"`
	Type        string `json:"type"`
	Archived 	bool `json:"archived"`
}

type user struct {
	Userid string
	Username string
	Password string

}

func GetAllTickets(session gocqlx.Session) error {

	var mySlice []string
	var query = session.Query("SELECT * FROM meed.ticket",mySlice)

    if rows, err := query.Iter().SliceMap(); err == nil {
        for _, row := range rows {
            fmt.Printf("%v\n", row)
        }
    } else {
        panic("Query error: " + err.Error())
    }
	return nil
}

func CreateTicketTable(session gocqlx.Session) error {
	session.ExecStmt(`DROP KEYSPACE meed`)

	err := session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS meed WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 2}`)
	if err != nil {
		fmt.Println("create keyspace:", err)
		return err
	}

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS meed.ticket (
		ticketid uuid PRIMARY KEY,
		title text,
		userid text,
		description text,
		reward text,
		lifespan int,
		type text,
		archived boolean
		)`)
	if err != nil {
		fmt.Println("create table:", err)
		return err
	}

	return nil
}

func SelectTicketByID(session gocqlx.Session, ticketTable *table.Table, uuid gocql.UUID) {
	ticketrows := ticketTable.SelectQuery(session)
	ticketrows.BindStruct(&ticket{
		Ticketid: uuid,
	})
	var tickets []ticket

	if err := ticketrows.Select(&tickets); err != nil {
		log.Fatal("Select() failed:", err)
	}
	fmt.Println(tickets)
	fmt.Println("selected correctly!")
}

func CreateTicket(session gocqlx.Session) {

	insertTicket := qb.Insert("meed.ticket").
	Columns("ticketid", 
			"title", 
			"userid", 
			"description", 
			"reward", 
			"lifespan", 
			"type", 
			"archived").Query(session)

	newTicket := ticket{}
	err := faker.FakeData(&newTicket)
	if err != nil {
		fmt.Println(err)
	}

	insertTicket.BindStruct(newTicket)
	// insertTicket.BindStruct(ticket{
	// 	Ticketid: uuid,
	// 	Title: "this is the insert query",
	// 	Userid: "johnny",
	// 	Description: "automatic insert",
	// 	Reward: "the joy of not having to rewrite this",
	// 	Lifespan: 5,
	// 	Type: "service",
	// 	Archived: false,
	// })

	if err := insertTicket.ExecRelease(); err != nil {
		log.Fatal("ExecRelease() failed:", err)
	}
}



func MustParseUUID(s string) gocql.UUID {
	u, err := gocql.ParseUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

func DeleteTicket(session *gocql.Session) error {

	uuid, merr := gocql.ParseUUID("c63e71f0-936e-11ea-bb37-0242ac130002")
	if merr != nil {
		log.Fatal(merr)
	}

	r := ticket {
		Ticketid: uuid,
	}

	w := qb.EqNamed("ticketid", uuid.String())
	fmt.Println(uuid.String())
	stmt, names := qb.Delete("mykeyspace.ticket").From("mykeyspace.ticket").Where(w).ToCql()
		fmt.Println(stmt, names)
        q := gocqlx.Query(session.Query(stmt), names).BindStruct(r)
		defer q.Release()

		fmt.Println(q)

        err := q.ExecRelease() 
		fmt.Println(err)
        if err != nil {
            log.Fatal(err)
        }
	return nil
}