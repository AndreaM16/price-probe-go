package itemfactory

import (
	"fmt"
	"net/http"
)

func ItemReceiver(r *http.Request) {

	key := GetParameterFromUrlByKey("key", r)
	switch key {
	case "id":
		fmt.Println("Got Id")
	case "pid":
		fmt.Println("Got Pid")
	case "category":
		fmt.Println("Got Catgory")
	case "title":
		fmt.Println("Got Title")
	case "url":
		fmt.Println("Got Url")
	}

}
