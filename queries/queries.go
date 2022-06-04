package queries

import (
	// "encoding/json"
	"fmt"
	"log"

	// "errors"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/databitio/go_server/datatypes"

	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

//perhaps future cqlx session object
type Gosess struct {
	Session gocqlx.Session
}

//*************************************************************************
//Metadata and create table used to filter data
func CreateTicketMetadata() table.Metadata {
	ticketMetadata := table.Metadata{
		Name:    "meed.ticket",
		Columns: []string{"ticketid", "title", "userid", "description", "reward", "lifespan", "type", "archived"},
		PartKey: []string{"ticketid"},
		SortKey: []string{"userid"},
	}
	return ticketMetadata
}

func CreateTable(metadata table.Metadata) *table.Table {
	return table.New(metadata)
}

//*************************************************************************
func GetByID(session gocqlx.Session, uuid gocql.UUID) (map[string]interface{}, error) {
	w := qb.EqNamed("ticketid", "")
	q := qb.Select("meed.ticket").Where(w).Query(session).Bind(uuid.String())

	rows, err := q.Iter().SliceMap()
	if err != nil {
		fmt.Println("Query error: " + err.Error())
	}
	return rows[0], err
}

func GetAllTickets(session gocqlx.Session) ([]map[string]interface{}, error) {

	q := qb.Select("meed.ticket").Query(session)

	rows, err := q.Iter().SliceMap()
	if err == nil {
		return rows, nil
	} else {
		panic("Query error: " + err.Error())
	}
}

func CreateNewServer(session gocqlx.Session) error {

	session.ExecStmt(`DROP KEYSPACE meed`)
	err := session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS meed WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 2}`)
	if err != nil {
		fmt.Println("create keyspace:", err)
		return err
	}

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS meed.ticket (
		ticketid uuid,
		userid uuid,
		serverid uuid,
		title text,
		description text,
		reward text,
		lifespan int,
		type text,
		archived boolean,
		PRIMARY KEY (ticketid, userid)
		)`)
	if err != nil {
		fmt.Println("create table:", err)
		return err
	}

	return nil
}

func SelectTicketByLocalID(session gocqlx.Session, ticketTable *table.Table, uuid gocql.UUID) {
	ticketrows := ticketTable.SelectQuery(session)
	ticketrows.BindStruct(&datatypes.Ticket{
		Ticketid: uuid,
	})
	var tickets []datatypes.Ticket

	if err := ticketrows.Select(&tickets); err != nil {
		log.Fatal("Select() failed:", err)
	}
	fmt.Println(tickets)
	fmt.Println("selected correctly!")
}

func CreateFakeTicket() datatypes.Ticket {

	newTicket := datatypes.Ticket{}
	err := faker.FakeData(&newTicket)
	if err != nil {
		fmt.Println(err)
	}
	return newTicket
}

func CreateTicket(session gocqlx.Session, newTicket datatypes.Ticket) {

	insertTicket := qb.Insert("meed.ticket").
		Columns("ticketid",
			"userid",
			"serverid",
			"title",
			"description",
			"reward",
			"lifespan",
			"type",
			"archived").Query(session)

	insertTicket.BindStruct(newTicket)

	if err := insertTicket.ExecRelease(); err != nil {
		log.Fatal("ExecRelease() failed:", err)
	}
}

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

func MustParseUUID(s string) gocql.UUID {
	u, err := gocql.ParseUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

//Couldn't get Delete cqlx to work; improvised with Delete statement
func DeleteTicket(session gocqlx.Session, uuid gocql.UUID) error {

	w := qb.EqNamed("ticketid", uuid.String())
	deleteTicket := qb.Delete("meed.ticket").Where(w).Query(session).Bind(uuid.String())

	if err := deleteTicket.ExecRelease(); err != nil {
		log.Fatal("ExecRelease() failed:", err)
		return err
	}

	fmt.Println("deleted successfully!")
	return nil
}
