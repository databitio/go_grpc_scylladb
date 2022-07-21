package api

import (
	"context"
	"fmt"
	"log"

	"github.com/databitio/go_server/queries"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"github.com/databitio/go_server/server_grpc/server/db"
	"github.com/databitio/go_server/server_grpc/server/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GoServer) UpdateTicket(ctx context.Context, in *pb.TicketInfo) (*pb.TicketID, error) {
	log.Println("TestInput was invoked")

	newTicket := utils.TicketMessageToTicket(in)

	err := queries.UpdateTicket(db.Session, newTicket)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Update ticket failed: %d\n", err),
		)
	}

	res := &pb.TicketID{
		Result: newTicket.Ticketid.String(),
	}

	return res, nil
}
