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

type GoServer struct {
	pb.TicketServiceServer
}

func (s *GoServer) CreateTicket(ctx context.Context, in *pb.TicketInfo) (*pb.TicketID, error) {
	log.Println("CreateTicket was invoked")

	newTicket := utils.TicketMessageToTicket(in)

	err := queries.CreateTicket(db.Session, newTicket)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Create ticket failed: %d\n", err),
		)
	}
	req := &pb.TicketID{
		Result: in.Ticketid,
	}

	return req, nil
}
