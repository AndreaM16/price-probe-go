package itemfactory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/andream16/price-probe-go/api"
	"github.com/andream16/price-probe-go/api/item/entity"
	"github.com/gocql/gocql"
)

// ItemsReceiver takes an *http.Request and a *gocql.Session
// returns []byte containing the json response of the query result
func ItemsReceiver(r *http.Request, s *gocql.Session) []byte {
	page, pageErr := takeParamFromURL(r, "page")
	if pageErr != nil {
		fmt.Println("Bad parameter for page. " + pageErr.Error())
		os.Exit(1)
	}
	size, sizeErr := takeParamFromURL(r, "size")
	if sizeErr != nil {
		fmt.Println("Bad parameter for size. " + sizeErr.Error())
		os.Exit(1)
	}
	requestBody := &api.RequestBody{page, size}
	items := getItemsFromCassandra(requestBody, s)
	return itemsResponseBuilder(items)
}

// itemsResponseBuilder takes items from the query and marshals them into json bytes
func itemsResponseBuilder(queryResult itementity.Items) []byte {
	var response []byte
	if len(queryResult.Items) == 0 {
		return response
	}
	response, _ = json.Marshal(queryResult.Items)
	return response
}

// getItemsFromCassandra takes a request body (key, value) and retrieves the items given by those parameters
func getItemsFromCassandra(requestBody *api.RequestBody, s *gocql.Session) itementity.Items {
	var item itementity.Item
	items := make([]itementity.Item, 0)
	iter := s.Query(`SELECT * FROM `+itemsTable+` LIMIT ?`, requestBody.Value).Consistency(gocql.One).Iter()
	for {
		row := map[string]interface{}{
			"item":        &item.ID,
			"category":    &item.Category,
			"description": &item.Description,
			"img":         &item.Img,
			"pid":         &item.Pid,
			"title":       &item.Title,
			"url":         &item.URL,
		}
		if !iter.MapScan(row) {
			break
		}
		if len(item.ID) > 0 {
			items = append(items, item)
		}
	}
	return itementity.Items{items}
}

// takeParamFromURL takes an http request and a key, returns the value of that key from url parameters
func takeParamFromURL(r *http.Request, key string) (int, error) {
	p, e := strconv.Atoi(GetParameterFromURLByKey(key, r))
	return p, e
}
