package api

import (
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
)

func checkIfParametersInUrl(r *http.Request) bool {
	parameters := r.URL.Query()
	if len(parameters) < 2 {
		return false
	}
	return true
}

func checkRequestType(r *http.Request) string {
	parameters := r.URL.Query()
	var requestType string
	for k, _ := range parameters {
		switch k {
		case "key":
			requestType = "query"
		case "page":
			requestType = "plain"
		}
	}
	return requestType
}

func ItemHandler(s gocql.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(checkRequestType(r) + " miao")
		key := r.URL.Query().Get("key")
		if len(key) == 0 {
			page := r.URL.Query().Get("page")
			if len(page) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			} else {
				size := r.URL.Query().Get("size")
				if len(size) == 0 {
					w.WriteHeader(http.StatusBadRequest)
					return
				} else {
					fmt.Println("Got Page")
				}
			}
		} else {
			value := r.URL.Query().Get("value")
			if len(value) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			switch key {
			case "id":
				fmt.Println("Got Id")
			case "pid":
				fmt.Println("Got Pid")
			case "category":
				fmt.Println("Got Catgory")
			case "title":
				fmt.Println("Got Title")
			}
		}

	}
}
