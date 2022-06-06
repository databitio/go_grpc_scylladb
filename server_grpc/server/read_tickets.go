package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
)



func (s *Server) ReadTickets(in *emptypb.Empty, stream pb.TicketService_ReadTicketsServer) error {
	log.Println("ReadTickets was invoked")

	tickets, err:= queries.GetAllTickets(session)

	if err != nil {
		log.Fatalf("Failed to get all tickets: %v\n", err)
	}

	for _, ticket := range tickets {
		newTicket := ticketToTicketMessage(&ticket)
		fmt.Println(newTicket)
		stream.Send(newTicket)
	}

	return nil
}