package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"google.golang.org/grpc/status"
)

func goDeleteTicket(c pb.TicketServiceClient, id string) error {
	fmt.Println("goDeleteTicket client was invoked")

	req := &pb.TicketID{
		Result: id,
	}

	_, err := c.DeleteTicket(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())
		} else {
			log.Fatalf("A non GRPC error: %v\n", err)
			return err
		}
	}
	return nil
}
