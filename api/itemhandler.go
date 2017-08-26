package api

import (
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
)

func ItemHandler(s gocql.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

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
