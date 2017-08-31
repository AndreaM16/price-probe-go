package itemfactory

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	items := getItemsFromCassandra(page, size, s)
	return itemsResponseBuilder(items)
}

func itemsResponseBuilder(queryResult itementity.Items) []byte {
	var response []byte
	if len(queryResult.Items) == 0 {
		return response
	}
	response, _ = json.Marshal(queryResult.Items)
	return response
}

func getItemsFromCassandra(page int, size int, s *gocql.Session) itementity.Items {
	var item itementity.Item
	items := make([]itementity.Item, 0)
	iter := s.Query(`SELECT * FROM `+itemsTable+` LIMIT ?`, size).Consistency(gocql.One).Iter()
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

func takeParamFromURL(r *http.Request, key string) (int, error) {
	p, e := strconv.Atoi(GetParameterFromURLByKey(key, r))
	return p, e
}
