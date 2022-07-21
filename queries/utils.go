package queries

import (
	"fmt"
	"time"

	"github.com/databitio/go_server/datatypes"

	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
)

func CreateFakeTicket() datatypes.Ticket {

	newTicket := datatypes.Ticket{}
	err := faker.FakeData(&newTicket)
	if err != nil {
		fmt.Println(err)
	}
	return newTicket
}

func MustParseUUID(s string) gocql.UUID {
	u, err := gocql.ParseUUID(s)
	if err != nil {
		panic(err)
	}
	return u
}

func DBToTicket(ticket map[string]interface{}) (datatypes.Ticket, error) {

	newTicket := &datatypes.Ticket{
		Ticketid:    ticket["ticketid"].(gocql.UUID),
		Userid:      ticket["userid"].(gocql.UUID),
		Serverid:    ticket["serverid"].(gocql.UUID),
		Title:       ticket["title"].(string),
		Description: ticket["description"].(string),
		Reward:      ticket["reward"].(string),
		Lifespan:    ticket["lifespan"].(time.Time),
		Type:        ticket["type"].(string),
		Archived:    ticket["archived"].(bool),
		Status:      ticket["status"].(string),
		Claimed:     ticket["claimed"].(bool),
	}

	return *newTicket, nil
}
