package main

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func UpdateTicket(ctx context.Context, in *pb.TicketInfo) (*emptypb.Empty, error) {
	fmt.Println("UpdateTicket was invoked")

	ticket := ticketMessageToTicket(in)
	err := queries.UpdateTicket(session, ticket)
	if err != nil {
		fmt.Printf("Update Ticket error: %v\n", err)
	}

	return nil, nil
}
