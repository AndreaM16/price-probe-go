package pricefactory

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/andream16/price-probe-go/api"
	"github.com/andream16/price-probe-go/api/price/entity"
	"github.com/gocql/gocql"
)

// PricesReceiver takes an *http.Request and a *gocql.Session
// returns []byte containing the json response of the query result
func PricesReceiver(r *http.Request, s *gocql.Session) []byte {
	key := takeParamFromURL(r, "key")
	if len(key) == 0 {
		fmt.Println("Bad parameter for key.")
		os.Exit(1)
	}
	value := takeParamFromURL(r, "value")
	if len(value) == 0 {
		fmt.Println("Bad parameter for value.")
		os.Exit(1)
	}
	requestBody := &api.RequestBody{key, value}
	prices := getPricesFromCassandraByKey(requestBody, s)
	return pricesResponseBuilder(prices)
}

// pricesResponseBuilder prices from query and build a json out of them in bytes
func pricesResponseBuilder(queryResult priceentity.Prices) []byte {
	var response []byte
	if len(queryResult.Prices) == 0 {
		return response
	}
	response, _ = json.Marshal(queryResult.Prices)
	return response
}

// getPricesFromCassandraByKey takes a requestBody(key, value) and finds all the prices about that key and value
func getPricesFromCassandraByKey(requestBody *api.RequestBody, s *gocql.Session) priceentity.Prices {
	var dateSlice time.Time
	prices := make([]priceentity.Price, 0)
	iter := s.Query(`SELECT * FROM `+PriceTable+` WHERE `+(requestBody.Key).(string)+` = ?`, requestBody.Value.(string)).Consistency(gocql.One).Iter()
	for {
		var price priceentity.Price
		row := map[string]interface{}{
			"item":      &price.Item,
			"date":      &dateSlice,
			"price":     &price.Price,
			"estimated": &price.Estimated,
		}
		if !iter.MapScan(row) {
			break
		}
		if len(price.Item) > 0 {
			price = *parsePriceFields(&price, &dateSlice)
			prices = append(prices, price)
		}
	}
	return priceentity.Prices{prices}
}

// parsePriceFields takes a price and a timeSlice and parsed the price as desired
func parsePriceFields(price *priceentity.Price, dateSlice *time.Time) *priceentity.Price {
	price.Price = toFixed(price.Price, 2)
	price.Estimated = toFixed(price.Estimated, 2)
	price.Date = append(price.Date, dateSlice.Day())
	price.Date = append(price.Date, int(dateSlice.Month()))
	price.Date = append(price.Date, dateSlice.Year())
	return price
}

// takeParamFromURL takes an http request and a key and returns the value given by that key from url
func takeParamFromURL(r *http.Request, key string) string {
	return GetParameterFromURLByKey(key, r)
}

// toFixes takes a float and a precision and returns that float fixes to a lenght of precision
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

// round takes a float and rounds it
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
