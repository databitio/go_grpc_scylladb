package main

import (
	"context"
	"fmt"
	"log"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func (s *Server) CreateTicket(ctx context.Context, in *pb.TicketInfo) (*pb.TicketID, error) {
	log.Println("TestInput was invoked")

	newTicket := ticketMessageToTicket(in)

	queries.CreateTicket(session, newTicket)

	fmt.Println("completed ticket creation!")
	fmt.Println("parameter:\n", in)

	// if err != nil {
	// 	log.Fatalf("Failed to get all tickets: %v\n", err)
	// }

	return nil, nil
}
