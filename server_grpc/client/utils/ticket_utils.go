package utils

import (
	"time"

	"github.com/databitio/go_server/datatypes"
	"github.com/databitio/go_server/queries"
	"github.com/databitio/go_server/server_grpc/proto"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"github.com/gocql/gocql"
)

func TicketToTicketMessage(data *datatypes.Ticket) *pb.TicketInfo {
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

func TicketMessageToTicket(data *pb.TicketInfo) *datatypes.Ticket {
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

func CreateFakeTicketMessage(uuid gocql.UUID) *proto.TicketInfo {
	time := time.Now()

	newTicket := datatypes.Ticket{
		Ticketid:    uuid,
		Serverid:    uuid,
		Userid:      uuid,
		Title:       "this is created ticket query",
		Description: "newly created ticket",
		Reward:      "newly created ticket",
		Lifespan:    time,
		Type:        "service",
		Archived:    false,
		Status:      "updated!",
		Claimed:     false,
	}

	return TicketToTicketMessage(&newTicket)
}
