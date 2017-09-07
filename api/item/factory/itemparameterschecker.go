package itemfactory

import (
	"net/http"
	"reflect"
	"strconv"
)

const itemsTable = "itemst"

// CheckIfParametersAreValid takes an http Request
// returns a tuple containing bool, string
// if true, the query is valid and its type (plain, query) is T._2
// if false, the query is not valid and T,_2 is empty
func CheckIfParametersAreValid(r *http.Request) (bool, string) {
	paramsInURL := parametersInURL(r)
	if paramsInURL {
		switch getRequestType(r) {
		case "query":
			if queryParametersvalid(r) {
				return true, "query"
			}
			return false, ""
		case "plain":
			if plainParametersValid(r) {
				return true, "plain"
			}
			return false, ""
		}
	}
	return false, ""
}

// parametersInURL check if there are at least 2 parameters in url
func parametersInURL(r *http.Request) bool {
	parameters := r.URL.Query()
	if len(parameters) < 2 {
		return false
	}
	return true
}

// getRequestType check wheter we have to return an item or a slice of items
func getRequestType(r *http.Request) string {
	parameters := r.URL.Query()
	var requestType string
	for k := range parameters {
		switch k {
		case "key":
			requestType = "query"
		case "page":
			requestType = "plain"
		}
	}
	return requestType
}

// plainParametersValid checks whether or not items query is valid
func plainParametersValid(r *http.Request) bool {
	key := GetParameterFromURLByKey("page", r)
	value := GetParameterFromURLByKey("size", r)
	_, err := strconv.Atoi(key)
	if err != nil {
		return false
	}
	_, e := strconv.Atoi(value)
	if e != nil {
		return false
	}
	return true
}

// queryParametersvalid checks wheter or not item query is valid
func queryParametersvalid(r *http.Request) bool {
	parameters := r.URL.Query()
	for k, v := range parameters {
		if reflect.TypeOf(k).Kind() != reflect.String || reflect.TypeOf(v).Kind() != reflect.Slice {
			return false
		}
	}
	return true
}

// GetParameterFromURLByKey takes key and an http Request,
// returns the value of that key into URL parameters
func GetParameterFromURLByKey(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}
