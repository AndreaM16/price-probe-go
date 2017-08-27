package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../api/item/rest"
	"../configuration"
	"github.com/gorilla/mux"
)

func InitServer(conf configuration.Configuration) {

	router := mux.NewRouter()
	router.HandleFunc("/item", itemrest.ItemHandler())
	http.Handle("/", router)

	port := strconv.Itoa(conf.Server.Port)

	fmt.Println("Started server at port :" + port + ". Now listening . . .")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", port}, ""), router))

}
