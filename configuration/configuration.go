package configuration

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Server struct {
		Host string `json:"Host"`
		Port int    `json:"Port"`
	} `json:"Server"`
	Cassandra struct {
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		Keyspace string `json:"Keyspace"`
	} `json:"Cassandra"`
}

func InitConfiguration() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println("error " + err.Error())
		os.Exit(1)
	}
	return configuration
}

func getFileName() string {
	filename := []string{"conf", ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
