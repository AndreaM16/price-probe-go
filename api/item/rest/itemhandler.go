package itemrest

import (
	"net/http"

	"github.com/andream16/price-probe-go/api/item/factory"
	"github.com/gocql/gocql"
)

func ItemHandler(session *gocql.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		paramsValid, queryType := itemfactory.CheckIfParametersAreValid(r)
		if !paramsValid {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch queryType {
		case "query":
			itemfactory.ItemReceiver(r)
		case "plain":
			itemfactory.ItemsReceiver(r, session)
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
