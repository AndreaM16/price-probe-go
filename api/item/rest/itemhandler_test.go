package itemrest

import (
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/andream16/price-probe-go/api/item/factory"
)

const BaseURL = "http://localhost:8000/item"
const ItemsURL = "?page=1&size=16"

func TestItemHandler(t *testing.T) {
	var client http.Client
	resp, err := client.Get(BaseURL + ItemsURL)
	if err != nil {
		t.Fatalf("Unable to get " + BaseURL + ItemsURL)
	}
	paramsValid, queryType := itemfactory.CheckIfParametersAreValid(resp.Request)
	if !paramsValid {
		t.Fatalf("Invalid parameters for " + ItemsURL)
	}
	if reflect.TypeOf(queryType).Kind() != reflect.String {
		t.Fatalf("Invalid type for " + ItemsURL)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Unable to get " + BaseURL + ItemsURL + ". Got http.StatusCode " + strconv.Itoa(resp.StatusCode))
	}
}
