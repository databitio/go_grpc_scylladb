package conn

import (
	"log"

	pb "github.com/databitio/go_server/server_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectGRPC() pb.TicketServiceClient {
	var addr string = "localhost:50051"

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	log.Println("connected successfully!")
	defer conn.Close()
	return pb.NewTicketServiceClient(conn)
}
