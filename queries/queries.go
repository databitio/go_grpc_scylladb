package queries

import (
	// "encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	// "errors"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type Ticket struct {
	Ticketid    gocql.UUID `json:"ticketid"`
	Userid      gocql.UUID `json:"userid"`
	Serverid    gocql.UUID `json:"serverid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reward      string     `json:"reward"`
	Lifespan    int        `json:"lifespan"`
	Type        string     `json:"type"`
	Archived    bool       `json:"archived"`
}

type Server struct {
	Serverid     gocql.UUID   `json:"serverid"`
	Users        []gocql.UUID `json:"users"` //userids
	Name         string       `json:"name"`
	Queuepages   []Queuepage  `json:"queuepages"` //queuepageid
	Date_created time.Time    `json:"date_created"`
}
type Queuepage struct {
	Queuepageid  int         `json:"queuepageid"`
	Queues       []Queue     `json:"queues"` //[queueids],
	Section_name string      `json:"section_name"`
	Name         string      `json:"name"`
	Messengers   []Messenger `json:"messengers"` //[messenger, messenger],
}
type Queue struct {
	Queueid int          `json:"queueid"`
	Name    string       `json:"name"`
	Tickets []gocql.UUID `json:"tickets"` //[ticketids]
}
type Messenger struct {
	Messengerid int          `json:"messengerid"`
	Name        string       `json:"name"`
	Messages    []gocql.UUID `json:"messages"` //[message, message]
}
type Message struct {
	Messageid gocql.UUID `json:"messageid"`
	Userid    gocql.UUID `json:"userid"`
	Content   string     `json:"content"`
	Date      time.Time  `json:"date"`
}

type User struct {
	Userid       gocql.UUID   `json:"userid"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Date_created time.Time    `json:"date_created"`
	Subscription Subscription `json:"subscription"`
}

type Subscription struct {
	Active       bool      `json:"active"`
	Type         string    `json:"type"`
	Renewal_date time.Time `json:"renewal_date"`
}

func GetAllTickets(session gocqlx.Session) error {

	var mySlice []string
	var query = session.Query("SELECT * FROM meed.ticket", mySlice)

	if rows, err := query.Iter().SliceMap(); err == nil {
		for _, row := range rows {
			fmt.Printf("%v\n", row)
		}
	} else {
		panic("Query error: " + err.Error())
	}
	return nil
}

func CreateNewServer(session gocqlx.Session) error {

	session.ExecStmt(`DROP KEYSPACE meed`)
	err := session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS meed WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 2}`)
	if err != nil {
		fmt.Println("create keyspace:", err)
		return err
	}

	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS meed.ticket (
		ticketid uuid PRIMARY KEY,
		serverid uuid,
		userid text,
		title text,
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
	ticketrows.BindStruct(&Ticket{
		Ticketid: uuid,
	})
	var tickets []Ticket

	if err := ticketrows.Select(&tickets); err != nil {
		log.Fatal("Select() failed:", err)
	}
	fmt.Println(tickets)
	fmt.Println("selected correctly!")
}

func CreateFakeTicket() Ticket {

	newTicket := Ticket{}
	err := faker.FakeData(&newTicket)
	if err != nil {
		fmt.Println(err)
	}
	return newTicket
}

func CreateTicket(session gocqlx.Session, newTicket Ticket) {

	insertTicket := qb.Insert("meed.ticket").
		Columns("ticketid",
			"serverid",
			"userid",
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
	deleteTicket := qb.Delete("meed.ticket").Where(w).Query(session)

	query := strings.ReplaceAll(deleteTicket.Statement(), "?", uuid.String())

	if err := session.ExecStmt(query); err != nil {
		log.Fatal("ExecRelease() failed:", err)
	}

	deleteTicket.Release()
	// fmt.Println(deleteTicket.Release())
	return nil
}
