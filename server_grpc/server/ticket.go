package main

import (
	"github.com/databitio/go_server/datatypes"
	pb "github.com/databitio/go_server/server_grpc/proto"
)

func ticketToTicketMessage(data *datatypes.Ticket) *pb.TicketInfo {
	return &pb.TicketInfo{
		Ticketid:    data.Ticketid.String(),
		Userid:      data.Userid.String(),
		Serverid:    data.Serverid.String(),
		Title:       data.Title,
		Description: data.Description,
		Reward:      data.Reward,
		Lifespan:    data.Lifespan.String(),
		Type:        data.Type,
		Archived:    data.Archived,
		Status:      data.Status,
		Claimed:     data.Claimed,
	}
}

