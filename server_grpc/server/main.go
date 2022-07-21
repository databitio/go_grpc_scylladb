package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/databitio/go_server/server_grpc/proto"
	"github.com/databitio/go_server/server_grpc/server/api"
	"github.com/databitio/go_server/server_grpc/server/db"
	"github.com/scylladb/gocqlx/v2"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var session gocqlx.Session = db.ConnectToCluster()

func main() {
	var session gocqlx.Session = db.ConnectToCluster()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen on server: ", addr)
	}
	fmt.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	server := api.GoServer{}

	pb.RegisterTicketServiceServer(s, &server)

	defer session.Close()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
