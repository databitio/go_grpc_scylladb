package main

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

// var tickets = []ticket{
// 	{Ticketid: "ticket1", Title: "ticket title 1", Username: "username1", Description: "this is the first ticket", Reward: "16 copper", Lifespan: "10 days", Type: "service"},
// 	{Ticketid: "ticket2", Title: "ticket title 2", Username: "username2", Description: "this is the second ticket", Reward: "16 copper", Lifespan: "9 days", Type: "service"},
// 	{Ticketid: "ticket3", Title: "ticket title 3", Username: "username3", Description: "this is the third ticket", Reward: "16 copper", Lifespan: "8 days", Type: "service"},
// }

// func ticketByID(c *gin.Context) {
// 	id := c.Param("ticketid")
// 	ticket, err := getTicketByID(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ticket not found"})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, ticket)
// }

// func getTicketByID(id string) (*ticket, error) {
// 	for index, ticket := range tickets {
// 		if ticket.Ticketid == id {
// 			return &tickets[index], nil
// 		}
// 	}
// 	return nil, errors.New("ticket not found")
// }

// func getTickets(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, tickets)
// }

// func createTicket(c *gin.Context) {
// 	var newTicket ticket

// 	if err := c.BindJSON(&newTicket); err != nil {
// 		return
// 	}

// 	tickets = append(tickets, newTicket)
// 	c.IndentedJSON(http.StatusCreated, newTicket)
// }

// func queryAll(session *gocql.Session) {
// 	var query = session.Query("SELECT * FROM mykeyspace.ticket")
// 	return query
// }

// func readAll(c *gin.Context) {

//     var query := queryAll()

// 	if rows, err := query.Iter().SliceMap(); err == nil {
// 		for _, row := range rows {
// 			c.IndentedJSON(http.StatusOK, row)
// 		}
// 	}
// 	return
// }

func main() {

	var cluster = gocql.NewCluster("52.3.213.119", "34.225.225.53", "54.82.189.29")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "HubOg2xDLqp6K0E"}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")

	var session, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic("Failed to connect to cluster")
	}

	defer session.Close()

	// uuid, _ := gocql.RandomUUID()
	// ticketMetadata := table.Metadata{
	// 	Name:    "meed.ticket",
	// 	Columns: []string{"ticketid", "title", "userid", "description", "reward", "lifespan", "type", "archived"},
	// 	PartKey: []string{"ticketid"},
	// 	SortKey: []string{"userid"},
	// }
	// var ticketTable = table.New(ticketMetadata)

	// createTicketTable(session)
	// createTicket(session)
	getAllTickets(session)
	// selectTicketByID(session, uuid)
}

// func getAllTickets(session *gocql.Session) []ticket {

// 	stmt, _ := qb.Select("mykeyspace.ticket").ToCql()
// 		var tickets []ticket
// 		err := gocqlx.Select(&tickets, session.Query(stmt))
//         if err != nil {
//             log.Fatal(err)
//         }
// 	return tickets
// }
// stmt, names := qb.Insert("mykeyspace.ticket").Columns("my_table_id", "word", "definition").ToCql()

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

func getAllTickets(session gocqlx.Session) error {

	var mySlice []string
	var query = session.Query("SELECT * FROM meed.ticket",mySlice)
	log.Println(query)

    if rows, err := query.Iter().SliceMap(); err == nil {
        for _, row := range rows {
            fmt.Printf("%v\n", row)
        }
    } else {
        panic("Query error: " + err.Error())
    }
	return nil
}

func createTicketTable(session gocqlx.Session) error {
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

func selectTicketByID(session gocqlx.Session, ticketTable *table.Table, uuid gocql.UUID) {

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

func createTicket(session gocqlx.Session) {

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



func mustParseUUID(s string) gocql.UUID {
	u, err := gocql.ParseUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

func deleteTicket(session *gocql.Session) error {

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