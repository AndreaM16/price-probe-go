package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/andream16/price-probe-go/api/item/rest"
	"github.com/andream16/price-probe-go/cassandra"
	"github.com/andream16/price-probe-go/configuration"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Setting up configuration . . .")
	conf := configuration.InitConfiguration()
	fmt.Println("Successfully got configuration!")

	fmt.Println("Initializing Cassandra connection . . .")
	s, err := cassandra.InitCassandraClient(&conf)
	if err != nil {
		fmt.Println("Unable to set up Cassandra client!")
		os.Exit(1)
	}
	fmt.Println("Successfully initialized Cassandra connection . . .")

	fmt.Println("Starting server . . .")
	router := mux.NewRouter()
	router.HandleFunc("/item", itemrest.ItemHandler(&s))
	http.Handle("/", router)

	port := strconv.Itoa(conf.Server.Port)

	fmt.Println("Started server at port :" + port + ". Now listening . . .")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", port}, ""), router))

}
