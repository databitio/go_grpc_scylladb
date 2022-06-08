package main

import (
	"context"
	"fmt"
	"log"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func (s *Server) DeleteTicket(ctx context.Context, in *pb.TicketID) (*pb.TicketID, error) {
	log.Println("TestInput was invoked")

	queries.DeleteTicket(session, queries.MustParseUUID(in.Result))

	fmt.Println("completed ticket creation!")
	fmt.Println("parameter:\n", in)

	// if err != nil {
	// 	log.Fatalf("Failed to get all tickets: %v\n", err)
	// }

	return nil, nil
}
