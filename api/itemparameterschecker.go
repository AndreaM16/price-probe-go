package api

import (
	"net/http"
	"reflect"
	"strconv"
)

func parametersInUrl(r *http.Request) bool {
	parameters := r.URL.Query()
	if len(parameters) < 2 {
		return false
	}
	return true
}

func getRequestType(r *http.Request) string {
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

func plainParametersValid(r *http.Request) bool {
	key := getParameterFromUrlByKey("page", r)
	value := getParameterFromUrlByKey("size", r)
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

func queryParametersvalid(r *http.Request) bool {
	parameters := r.URL.Query()
	for k, v := range parameters {
		if reflect.TypeOf(k).Kind() != reflect.String || reflect.TypeOf(v).Kind() != reflect.String {
			return false
		}
	}
	return true
}

func getParameterFromUrlByKey(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

func checkIfParametersAreValid(r *http.Request) bool {
	paramsInUrl := parametersInUrl(r)
	if paramsInUrl {
		switch getRequestType(r) {
		case "query":
			if queryParametersvalid(r) {
				return true
			} else {
				return false
			}
		case "plain":
			if plainParametersValid(r) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
