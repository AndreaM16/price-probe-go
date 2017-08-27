package api

import (
	"fmt"
	"net/http"
)

func ItemHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		paramsValid, queryType := CheckIfParametersAreValid(r)
		if !paramsValid {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch queryType {
		case "query":
			key := GetParameterFromUrlByKey("key", r)
			switch key {
			case "id":
				fmt.Println("Got Id")
			case "pid":
				fmt.Println("Got Pid")
			case "category":
				fmt.Println("Got Catgory")
			case "title":
				fmt.Println("Got Title")
			case "url":
				fmt.Println("Got Url")
			}
		case "plain":
			fmt.Println("Got Page & Size")
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
