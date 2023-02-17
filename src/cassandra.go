package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Session *gocql.Session

func init() {
	err := godotenv.Load()
  	if err != nil {
    log.Println("Couldn't find .env file. Using local environment variables instead.")
  	}
	cluster := gocql.NewCluster(os.Getenv("DB_HOST"))
	cluster.Keyspace = os.Getenv("DB_KEYSPACE")
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	fmt.Println("cassandra connection initialised")
}
