package main

import (
	"context"
	"log"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func (s *Server) GetTicket(ctx context.Context, in *pb.TicketID) (*pb.TicketInfo, error) {
	log.Printf("Greet function was invoked with %v\n", in)

	uuid := queries.MustParseUUID(in.Result)
	selectedTicket, err := queries.GetByID(session, uuid)

	if err != nil {
		log.Fatalf("GetTicket failed: %v\n", err)
	}

	ticketInfo := ticketToTicketMessage(&selectedTicket)
	return ticketInfo, nil
}
