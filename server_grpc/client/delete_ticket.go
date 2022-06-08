package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func goDeleteTicket(c pb.TicketServiceClient, id string) {
	fmt.Println("goDeleteTicket client was invoked")

	req := &pb.TicketID{
		Result: id,
	}

	res, err := c.DeleteTicket(context.Background(), req)
	fmt.Println(res)
	if err != nil {
		log.Fatalf("goTestInput client failed: %v\n", err)
	}
}
