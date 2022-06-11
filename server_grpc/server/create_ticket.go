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

func (s *Server) CreateTicket(ctx context.Context, in *pb.TicketInfo) (*pb.TicketID, error) {
	log.Println("CreateTicket was invoked")

	newTicket := ticketMessageToTicket(in)

	err := queries.CreateTicket(session, newTicket)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Create ticket failed: %d\n", err),
		)
	}
	req := &pb.TicketID{
		Result: in.Ticketid,
	}

	return req, nil
}
