package db

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var Session gocqlx.Session

func ConnectToCluster() gocqlx.Session {
	var cluster = gocql.NewCluster("3.233.176.20", "54.208.199.255", "44.206.172.83")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "w7lTdugqIh1F2RC"}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")

	var session, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		panic("Failed to connect to cluster")
	}
	fmt.Println("Connected to DB!")
	Session = session
	return session
}
