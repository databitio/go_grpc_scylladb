package main

import (
	"github.com/databitio/go_server/server_grpc/client/conn"
	"github.com/databitio/go_server/server_grpc/client/test"
)

func main() {
	c := conn.ConnectGRPC()
	test.Endpoints(c)
}
