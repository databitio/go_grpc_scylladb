package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"google.golang.org/grpc/status"
)

func goGetTicket(c pb.TicketServiceClient, id string) *pb.TicketInfo {
	fmt.Println("getTicket client was invoked")

	req := &pb.TicketID{
		Result: id,
	}
	res, err := c.GetTicket(context.Background(), req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
		} else {
			log.Fatalf("A non GRPC error: %v\n", err)
			return nil
		}
	}
	return res
}
