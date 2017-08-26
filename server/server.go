package server

import (
	"fmt"
	"log"
	"net/http"

	"../api"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func InitServer(cassandraSession gocql.Session) {

	router := mux.NewRouter()
	router.HandleFunc("/item", api.ItemHandler(cassandraSession))
	http.Handle("/", router)

	fmt.Println("Listening on port :8000 . . .")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", router))

}
