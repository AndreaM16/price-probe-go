package pricerest

import (
	"net/http"

	"github.com/andream16/price-probe-go/api/price/factory"
	"github.com/gocql/gocql"
)

// PriceHandler takes a gocql.Session, an http.ResponseWriter and a http.Request
// return http.Response containing http code and json of the result
func PriceHandler(session *gocql.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		pricesResponse := pricefactory.PricesReceiver(r, session)
		w.WriteHeader(http.StatusOK)
		w.Write(pricesResponse)
		return
	}
}
