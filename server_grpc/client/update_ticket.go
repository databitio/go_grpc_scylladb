package main

import (
	"context"
	"fmt"

	pb "github.com/databitio/go_server/server_grpc/proto"
)

func goUpdateTicket(c pb.TicketServiceClient, in *pb.TicketInfo) {
	fmt.Println("updateTicket client was invoked")

	fmt.Println(context.Background())
	fmt.Println(in)
	_, err := c.UpdateTicket(context.Background(), in)

	if err != nil {
		fmt.Printf("GetTicket client failed: %v\n", err)
	}

}
