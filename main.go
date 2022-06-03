package main

import (
	// "encoding/json"
	// "fmt"
	// "log"
	"example/go_server/queries"
	"fmt"

	// "errors"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	// "github.com/scylladb/gocqlx/table"
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

	ticketid, _ := gocql.RandomUUID()
	// serverid, _ := gocql.RandomUUID()
	// ticketMetadata := table.Metadata{
	// 	Name:    "meed.ticket",
	// 	Columns: []string{"ticketid", "title", "userid", "description", "reward", "lifespan", "type", "archived"},
	// 	PartKey: []string{"ticketid"},
	// 	SortKey: []string{"userid"},
	// }

	// var ticketTable = table.New(ticketMetadata)

	// newTicket := queries.Ticket{
	// 		Ticketid: ticketid,
	// 		Serverid: serverid,
	// 		Userid: "johnny",
	// 		Title: "this is the insert query",
	// 		Description: "automatic insert",
	// 		Reward: "the joy of not having to rewrite this",
	// 		Lifespan: 5,
	// 		Type: "service",
	// 		Archived: false,
	// 		Subtitles: [],
	// 		Icons: [],
	// 	}
	
	// 	queries.Server{
	// 		Serverid: serverid,
	// 		Users: [userids],
	// 		Name: ,
	// 		Queuepages: [queuepageid],
	// 		date_created: ,

	// 	}
	// 	Queuepage {
	// 		Queuepageid: ,
	// 		Queues: [queueids],
	// 		sectionname: ,
	// 		queuepagename: ,
	// 		messengers: [messenger, messenger],
	// 	}
	// 	queries.Queue{
	// 		Queueid: ,
	// 		Name: ,
	// 		Tickets: [ticketids]
	// 	}
	// 	queries.messenger{
	// 		messengerid: ,
	// 		messages: [message, message],
	// 	}
	// 	queries.message{
	// 		messageid: ,
	// 		userid: ,
	// 		time: ,
	// 	}

	queries.CreateNewServer(session)
	// queries.CreateTicket(session, newTicket)
	queries.CreateTicket(session, queries.CreateFakeTicket())
	fmt.Println("added")
	queries.GetAllTickets(session)
	fmt.Println("subtracting...")
	queries.DeleteTicket(session, ticketid)
	queries.GetAllTickets(session)
	// selectTicketByID(session, uuid)
}
