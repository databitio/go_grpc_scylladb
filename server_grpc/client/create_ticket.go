package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func goCreateTicket(c pb.TicketServiceClient, in *pb.TicketInfo) {
	fmt.Println("goCreateTicket client was invoked")

	res, err := c.CreateTicket(context.Background(), in)
	fmt.Println(res)
	if err != nil {
		log.Fatalf("goTestInput client failed: %v\n", err)
	}
}
