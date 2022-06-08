package queries

import (
	// "encoding/json"
	"fmt"
	"log"
	"time"

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

func DBToTicket(ticket map[string]interface{}) (datatypes.Ticket, error) {

	newTicket := &datatypes.Ticket{
		Ticketid:    ticket["ticketid"].(gocql.UUID),
		Userid:      ticket["userid"].(gocql.UUID),
		Serverid:    ticket["serverid"].(gocql.UUID),
		Title:       ticket["title"].(string),
		Description: ticket["description"].(string),
		Reward:      ticket["reward"].(string),
		Lifespan:    ticket["lifespan"].(time.Time),
		Type:        ticket["type"].(string),
		Archived:    ticket["archived"].(bool),
		Status:      ticket["status"].(string),
		Claimed:     ticket["claimed"].(bool),
	}

	return *newTicket, nil
}

//*************************************************************************
func GetByID(session gocqlx.Session, uuid gocql.UUID) (datatypes.Ticket, error) {
	w := qb.EqNamed("ticketid", "")
	q := qb.Select("meed.ticket").Where(w).Query(session).Bind(uuid.String())

	rows, err := q.Iter().SliceMap()
	if err != nil {
		fmt.Println("Query error: " + err.Error())
	}
	ticket, err := DBToTicket(rows[0])
	return ticket, err
}

func UpdateTicket(session gocqlx.Session, ticket *datatypes.Ticket) error {
	// w := qb.EqNamed("ticketid", "")
	// x := qb.EqNamed("serverid", "")
	// y := qb.EqNamed("userid", "")
	q := qb.Update("meed.ticket").
		Set(
			"title",
			"description",
			"reward",
			"lifespan",
			"type",
			"archived",
			"status",
			"claimed").
		Where(qb.Eq("ticketid"), qb.Eq("userid")).Existing().
		Query(session).
		Bind(
			ticket.Title,
			ticket.Description,
			ticket.Reward,
			ticket.Lifespan,
			ticket.Type,
			ticket.Archived,
			ticket.Status,
			ticket.Claimed,
			ticket.Ticketid,
			// ticket.Serverid,
			ticket.Userid,
		)

	defer q.Release()
	fmt.Println(q)

	if err := q.Exec(); err != nil {
		log.Fatalf("Error updating ticket: %v\n", err)
		return err
	}

	fmt.Println("Updated successfully!")
	return nil
}

func GetAllTickets(session gocqlx.Session) ([]datatypes.Ticket, error) {

	q := qb.Select("meed.ticket").Query(session)

	defer q.Release()

	var tickets []datatypes.Ticket
	if rows, err := q.Iter().SliceMap(); err == nil {
		for _, row := range rows {
			ticket, err := DBToTicket(row)
			if err == nil {
				tickets = append(tickets, ticket)
			}
		}
	}
	return tickets, nil
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
		lifespan date,
		type text,
		archived boolean,
		status text,
		claimed boolean,
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

func CreateTicket(session gocqlx.Session, newTicket *datatypes.Ticket) {

	insertTicket := qb.Insert("meed.ticket").
		Columns("ticketid",
			"userid",
			"serverid",
			"title",
			"description",
			"reward",
			"lifespan",
			"type",
			"archived",
			"status",
			"claimed").Query(session)

	insertTicket.BindStruct(*newTicket)

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
