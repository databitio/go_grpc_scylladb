package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func goGetTicket(c pb.TicketServiceClient, id string) *pb.TicketInfo {
	fmt.Println("getTicket client was invoked")

	req := &pb.TicketID{
		Result: id,
	}
	res, err := c.GetTicket(context.Background(), req)
	if err != nil {
		log.Fatalf("GetTicket client failed: %v\n", err)
	}
	return res
}
