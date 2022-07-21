package api

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"github.com/databitio/go_server/server_grpc/server/db"
	"github.com/databitio/go_server/server_grpc/server/utils"
)

func (s *GoServer) ReadTickets(in *emptypb.Empty, stream pb.TicketService_ReadTicketsServer) error {
	log.Println("ReadTickets was invoked")

	tickets, err := queries.GetAllTickets(db.Session)
	log.Println("Exited query")

	if err != nil {
		return status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Update ticket failed: %d\n", err),
		)
	}

	for _, ticket := range tickets {
		newTicket := utils.TicketToTicketMessage(&ticket)
		stream.Send(newTicket)
	}

	return nil
}
