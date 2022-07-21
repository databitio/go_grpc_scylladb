package test

import (
	"log"

	"github.com/databitio/go_server/server_grpc/client/api"
	"github.com/databitio/go_server/server_grpc/client/utils"
	pb "github.com/databitio/go_server/server_grpc/proto"
	"github.com/gocql/gocql"
)

type GoClient struct {
	c pb.TicketServiceClient
}

func Endpoints(c pb.TicketServiceClient) {

	uuid, _ := gocql.RandomUUID()
	ticket := utils.CreateFakeTicketMessage(uuid)

	client := &GoClient{
		c: c,
	}

	client.testCreateTicket(ticket)
	client.testGetTicket(uuid)
	client.testUpdateTicket(ticket)
	client.testGetTicket(uuid)
	client.testDeleteTicket(uuid)
	client.testReadTickets()

	log.Println("All operations complete! Shutting down...")
}

func (client GoClient) testCreateTicket(ticket *pb.TicketInfo) {
	log.Println("Creating ticket...")
	err := api.GoCreateTicket(client.c, ticket)
	if err != nil {
		log.Printf("Error creating ticket: %v\n", err)
		return
	}
	log.Println("Successfully added ticket!")
}

func (client GoClient) testGetTicket(uuid gocql.UUID) {
	log.Println("Getting ticket by id...")
	ticket := api.GoGetTicket(client.c, uuid.String())
	log.Printf("Ticket found!: %v\n", ticket)
}

func (client GoClient) testUpdateTicket(ticket *pb.TicketInfo) {
	ticket.Title = "Updated title!"
	log.Println("Updating this tickets title...")
	_, err := api.GoUpdateTicket(client.c, ticket)
	if err != nil {
		log.Printf("Error updating ticket: %v\n", err)
		return
	}
	log.Println("Ticket updated successfully!")
}

func (client GoClient) testDeleteTicket(uuid gocql.UUID) {
	log.Println("Deleting this ticket from db...")
	err := api.GoDeleteTicket(client.c, uuid.String())
	if err != nil {
		log.Printf("Error deleting ticket: %v\n", err)
		return
	}
	log.Println("Delete success!")
}

func (client GoClient) testReadTickets() {
	log.Println("Printing out all tickets...")
	alltickets := api.GoReadTickets(client.c)
	for index, ticket := range alltickets {
		log.Println(index, utils.TicketMessageToTicket(ticket))
	}
}
