package main

import (
	"context"
	"fmt"
	"log"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTicket(ctx context.Context, in *pb.TicketID) (*pb.TicketInfo, error) {
	log.Printf("Get ticket was invoked with %v\n", in)

	uuid := queries.MustParseUUID(in.Result)
	selectedTicket, err := queries.GetByID(session, uuid)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Update ticket failed: %d\n", err),
		)
	}

	ticketInfo := ticketToTicketMessage(&selectedTicket)
	return ticketInfo, nil
}
