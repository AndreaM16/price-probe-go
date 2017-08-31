package itemfactory

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andream16/price-probe-go/api/item/entity"
	"github.com/gocql/gocql"
)

// ItemReceiver takes http.Request and a gocql.Session
// returns []byte containing json of the query result
func ItemReceiver(r *http.Request, s *gocql.Session) []byte {
	key := GetParameterFromURLByKey("key", r)
	value := GetParameterFromURLByKey("value", r)
	item := getItemFromCassandraByKey(key, value, s)
	return itemResponseBuilder(item)
}

func itemResponseBuilder(queryResult itementity.Item) []byte {
	var response []byte
	response, _ = json.Marshal(queryResult)
	return response
}

func getItemFromCassandraByKey(key string, value interface{}, s *gocql.Session) itementity.Item {
	var item itementity.Item
	if err := s.Query(`SELECT * FROM `+itemsTable+` WHERE `+key+` = ? ALLOW FILTERING`, value).Consistency(gocql.One).Scan(&item.ID, &item.Category, &item.Description, &item.Img, &item.Pid, &item.Title, &item.URL); err != nil {
		log.Fatal(err)
	}
	return item
}
