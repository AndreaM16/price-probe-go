package itemrest

import (
	"net/http"

	"../factory"
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
			itemfactory.ItemReceiver(r)
		case "plain":
			itemfactory.ItemsReceiver(r)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
