package main

import (
	"fmt"
	"log"
	"time"

	"github.com/databitio/go_server/datatypes"
	"github.com/gocql/gocql"

	// "github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	log.Println("connected successfully!")
	defer conn.Close()
	c := pb.NewTicketServiceClient(conn)

	uuid, _ := gocql.RandomUUID()
	time := time.Now()

	newTicket := datatypes.Ticket{
		Ticketid:    uuid,
		Serverid:    uuid,
		Userid:      uuid,
		Title:       "this is created ticket query",
		Description: "newly created ticket",
		Reward:      "newly created ticket",
		Lifespan:    time,
		Type:        "service",
		Archived:    false,
		Status:      "updated!",
		Claimed:     false,
	}

	ticketinfo := ticketToTicketMessage(&newTicket)


	fmt.Println("Creating ticket...")
	err = goCreateTicket(c, ticketinfo)
	if err != nil {
		fmt.Printf("Error creating ticket: %v\n", err)
	}
	fmt.Println("Successfully added ticket!")

	fmt.Println("Getting ticket by id...")
	myticket := goGetTicket(c, uuid.String())
	fmt.Printf("Ticket found!: %v\n", myticket)

	newTicket.Title = "Updated title!"
	fmt.Println("Updating this tickets title...")
	ticketinfo = ticketToTicketMessage(&newTicket)
	_, err = goUpdateTicket(c, ticketinfo)
	if err != nil {
		fmt.Printf("Error updating ticket: %v\n", err)
	}
	fmt.Println("Ticket updated successfully!")
	
	fmt.Println("Getting updated ticket...")
	myticket = goGetTicket(c, uuid.String())
	fmt.Printf("Ticket found!: %v\n", myticket)

	fmt.Println("Deleting this ticket from db...")
	err = goDeleteTicket(c, uuid.String())
	if err != nil {
		fmt.Printf("Error deleting ticket: %v\n", err)
	}
	fmt.Println("Delete success! Printing out all tickets...")
	
	alltickets := readTickets(c)
	for index, ticket := range alltickets {
		fmt.Println(index, *ticketMessageToTicket(ticket))
	}

	fmt.Println("All operations complete! Shutting down...")
}

//Testing all CRUD operations
	// uuid := queries.MustParseUUID("44573233-4c12-1d06-2c63-0910604a1816")
	// time, _ := time.Parse("2006-01-02T15:04:05.000Z", "2232-12-23 00:00:00 +0000 UTC")

		// ticket := &pb.TicketInfo{
	// 	Ticketid:    "44573233-4c12-1d06-2c63-0910604a1816",
	// 	Userid:      "44573233-4c12-1d06-2c63-0910604a1816",
	// 	Serverid:    "44573233-4c12-1d06-2c63-0910604a1816",
	// 	Title:       "This is the updated ticket!",
	// 	Description: "updated ticket description",
	// 	Reward:      "updated ticket reward",
	// 	Lifespan:    "2296-03-25",
	// 	Type:        "updated ticket type",
	// 	Archived:    false,
	// 	Status:      "updated!",
	// 	Claimed:     false,
	// }
	// fmt.Println("ticket made!")

	// newTicket := datatypes.Ticket{
	// 	Ticketid:    uuid,
	// 	Serverid:    uuid,
	// 	Userid:      uuid,
	// 	Title:       "this is created ticket query",
	// 	Description: "newly created ticket",
	// 	Reward:      "newly created ticket",
	// 	Lifespan:    time,
	// 	Type:        "service",
	// 	Archived:    false,
	// 	Status:      "updated!",
	// 	Claimed:     false,
	// }

	// ticketinfo := ticketToTicketMessage(&newTicket)

	// fmt.Println("Reading all tickets...")
	// alltickets := readTickets(c)
	// for index, ticket := range alltickets {
	// 	fmt.Println(index, *ticketMessageToTicket(ticket))
	// }
	// fmt.Println("Getting ticket by id...")
	// myticket := goGetTicket(c, "44573233-4c12-1d06-2c63-0910604a1816")
	// fmt.Printf("Ticket found!: %v\n", myticket)
	// newTicket.Title = "this is the updated title! hooray!"
	// fmt.Println("Updating this tickets title...")
	// ticketinfo = ticketToTicketMessage(&newTicket)
	// _, err = goUpdateTicket(c, ticketinfo)
	// if err != nil {
	// 	fmt.Printf("Error updating ticket: %v\n", err)
	// }
	// fmt.Println("Ticket updated successfully!")
	// myticket = goGetTicket(c, "44573233-4c12-1d06-2c63-0910604a1816")
	// fmt.Printf("New ticket: %v\n", myticket)
	// fmt.Println("Deleting this ticket from db...")
	// err = goDeleteTicket(c, "44573233-4c12-1d06-2c63-0910604a1816")
	// if err != nil {
	// 	fmt.Printf("Error deleting ticket: %v\n", err)
	// }
	// fmt.Println("Delete success! Printing out all tickets...")
	// alltickets = readTickets(c)
	// for index, ticket := range alltickets {
	// 	fmt.Println(index, *ticketMessageToTicket(ticket))
	// }
	// fmt.Println("Adding ticket back to database...")
	// err = goCreateTicket(c, ticketinfo)
	// if err != nil {
	// 	fmt.Printf("Creating deleting ticket: %v\n", err)
	// }
	// fmt.Println("Successfully added ticket! Shutting down...")