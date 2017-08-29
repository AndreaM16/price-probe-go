package itemfactory

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/andream16/price-probe-go/api/item/entity"

	"github.com/gocql/gocql"
)

func ItemsReceiver(r *http.Request, s *gocql.Session) {
	page, pageErr := takeParamFromUrl(r, "page")
	if pageErr != nil {
		fmt.Println("Bad parameter for page. " + pageErr.Error())
		os.Exit(1)
	}
	size, sizeErr := takeParamFromUrl(r, "size")
	if sizeErr != nil {
		fmt.Println("Bad parameter for size. " + sizeErr.Error())
		os.Exit(1)
	}
	getItemsFromCassandra(page, size, s)
}

// &item.ID, &item.Category, &item.Description, &item.Img, &item.Pid, &item.Title, &item.URL
func getItemsFromCassandra(page int, size int, s *gocql.Session) {
	var item itementity.Item
	// var item itementity.Item
	iter := s.Query(`SELECT * FROM itemst LIMIT ?`, size).Consistency(gocql.One).Iter()
	for {
		// New map each iteration
		row := map[string]interface{}{
			"item":     &item.ID,
			"category": &item.Category,
		}
		if !iter.MapScan(row) {
			break
		}
		// Do things with row
		fmt.Printf("Id: %s Category: %s\n", item.ID, item.Category)
	}
}

func takeParamFromUrl(r *http.Request, key string) (int, error) {
	p, e := strconv.Atoi(GetParameterFromUrlByKey(key, r))
	return p, e
}
