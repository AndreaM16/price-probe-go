package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../api/item"
	"../configuration"
	"github.com/gorilla/mux"
)

func InitServer(conf configuration.Configuration) {

	router := mux.NewRouter()
	router.HandleFunc("/item", api.ItemHandler())
	http.Handle("/", router)

	fmt.Println(conf.Server)
	port := strconv.Itoa(conf.Server.Port)
	fmt.Println(port)

	fmt.Println(strings.Join([]string{"Listening on port ", port, " . . ."}, " "))
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(strings.Join([]string{":", port}, ""), router))

}
