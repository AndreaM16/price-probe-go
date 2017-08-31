package pricefactory

import "net/http"

const PriceTable = "pricest"

// GetParameterFromURLByKey takes key and an http Request,
// returns the value of that key into URL parameters
func GetParameterFromURLByKey(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}
