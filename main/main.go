package main

import (
	"fmt"
	"log"

	"../cassandramanager"
	"../server"
)

func main() {

	fmt.Println("Trying to enstablish a connection to a Cassandra Cluster . . .")
	s, err := cassandramanager.InitCassandraClient()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Connection to Cassandra Cluster successfully enstabilshed!")
	}

	fmt.Println("Starting server . . .")
	server.InitServer(s)

}
