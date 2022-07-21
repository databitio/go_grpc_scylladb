package api

import (
	"context"
	"fmt"
	"log"

	// "github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"google.golang.org/grpc/status"
)

func GoCreateTicket(c pb.TicketServiceClient, in *pb.TicketInfo) error {
	fmt.Println("goCreateTicket client was invoked")

	_, err := c.CreateTicket(context.Background(), in)
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
