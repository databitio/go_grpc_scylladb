package main

import (
	"time"

	"github.com/databitio/go_server/datatypes"
	"github.com/databitio/go_server/queries"
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

func ticketMessageToTicket(data *pb.TicketInfo) *datatypes.Ticket {
	time, _ := time.Parse("2006-01-02T15:04:05.000Z", data.Lifespan)
	return &datatypes.Ticket{
		Ticketid:    queries.MustParseUUID(data.Ticketid),
		Userid:      queries.MustParseUUID(data.Userid),
		Serverid:    queries.MustParseUUID(data.Serverid),
		Title:       data.Title,
		Description: data.Description,
		Reward:      data.Reward,
		Lifespan:    time,
		Type:        data.Type,
		Archived:    data.Archived,
		Status:      data.Status,
		Claimed:     data.Claimed,
	}
}
