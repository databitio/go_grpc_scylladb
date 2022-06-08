package main

import (
	"fmt"
	"log"
	"time"

	"github.com/databitio/go_server/datatypes"
	"github.com/databitio/go_server/queries"
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
	uuid := queries.MustParseUUID("44573233-4c12-1d06-2c63-0910604a1816")
	time, _ := time.Parse("2006-01-02T15:04:05.000Z", "2232-12-23 00:00:00 +0000 UTC")

	newTicket := datatypes.Ticket{
		Ticketid:    uuid,
		Serverid:    uuid,
		Userid:      uuid,
		Title:       "this is the insert query",
		Description: "automatic insert",
		Reward:      "the joy of not having to rewrite this",
		Lifespan:    time,
		Type:        "service",
		Archived:    false,
		Status:      "updated!",
		Claimed:     false,
	}

	fmt.Println("ticket conversion begin...")
	ticketinfo := ticketToTicketMessage(&newTicket)
	fmt.Println(ticketinfo)
	fmt.Println(c)

	// myticket := goGetTicket(c, "44573233-4c12-1d06-2c63-0910604a1816")
	// goUpdateTicket(c, ticketinfo)
	goUpdateTicket(c, ticketinfo)
	// readTickets(c)
	// fmt.Println(myticket)
}
