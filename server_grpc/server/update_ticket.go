package main

import (
	"context"
	"fmt"
	"log"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func (s *Server) UpdateTicket(ctx context.Context, in *pb.TicketInfo) (*pb.TicketID, error) {
	log.Println("TestInput was invoked")

	newTicket := ticketMessageToTicket(in)

	err := queries.UpdateTicket(session, newTicket)

	fmt.Println("completed update!")
	fmt.Println("parameter:\n", in)

	if err != nil {
		log.Fatalf("Failed to get all tickets: %v\n", err)
	}

	return nil, nil
}
