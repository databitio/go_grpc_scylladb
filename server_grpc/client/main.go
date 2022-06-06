package main

import (
	"log"

	pb "github.com/databitio/go_server/server_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	log.Println("connected successfully!")
	defer conn.Close()
	c := pb.NewTicketServiceClient(conn)

	newTicket := goGetTicket(c, "44573233-4c12-1d06-2c63-0910604a1816")
	log.Println(newTicket)
}
