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

func (s *Server) DeleteTicket(ctx context.Context, in *pb.TicketID) (*pb.TicketID, error) {
	log.Println("TestInput was invoked")

	err := queries.DeleteTicket(session, queries.MustParseUUID(in.Result))

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Delete ticket failed: %d\n", err),
		)
	}

	return in, nil
}
