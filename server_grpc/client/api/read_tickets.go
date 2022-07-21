package api

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/databitio/go_server/server_grpc/proto"
)

func GoReadTickets(c pb.TicketServiceClient) []*pb.TicketInfo {
	fmt.Println("readTickets client was invoked")

	var allTickets []*pb.TicketInfo
	stream, err := c.ReadTickets(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("ReadTickets client failed: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Something happened: %v\n", err)
			continue
		}

		allTickets = append(allTickets, res)
	}
	return allTickets
}
