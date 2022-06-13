package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/databitio/go_server/server_grpc/proto"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"google.golang.org/grpc"
)


func ConnectToCluster() gocqlx.Session {
	var cluster = gocql.NewCluster("3.233.176.20", "54.208.199.255", "44.206.172.83")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "w7lTdugqIh1F2RC"}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")

	var session, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic("Failed to connect to cluster")
	}
	fmt.Println("Connected to DB!")
	return session
}

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.TicketServiceServer
}

var session gocqlx.Session = ConnectToCluster()

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen on server: ", addr)
	}

	fmt.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterTicketServiceServer(s, &Server{})

	defer session.Close()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
