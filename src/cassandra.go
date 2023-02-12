package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"os"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster(os.Getenv("DB_HOST"))
	cluster.Keyspace = os.Getenv("DB_KEYSPACE")
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	fmt.Println("cassandra connection initialised")
}
