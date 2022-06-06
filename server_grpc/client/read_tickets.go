package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/databitio/go_server/server_grpc/proto"
)

func readTickets(c pb.TicketServiceClient) {
	fmt.Println("readTickets client was invoked")

	stream, err := c.ReadTickets(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("ReadTickets client failed: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil{
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}