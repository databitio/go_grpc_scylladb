package queries

import (
	// "encoding/json"
	"fmt"
	"time"

	// "errors"
	// "net/http"

	// "github.com/gin-gonic/gin"
	"github.com/databitio/go_server/datatypes"

	"github.com/bxcodec/faker/v3"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
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

//Metadata and create table used to filter data
func CreateTicketMetadata() table.Metadata {
	ticketMetadata := table.Metadata{
		Name:    "meed.ticket",
		Columns: []string{"ticketid", "title", "userid", "description", "reward", "lifespan", "type", "archived"},
		PartKey: []string{"ticketid"},
		SortKey: []string{"userid"},
	}
	return ticketMetadata
}

func CreateTable(metadata table.Metadata) *table.Table {
	return table.New(metadata)
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