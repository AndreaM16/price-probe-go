package main

import (
	"testing"

	"github.com/andream16/price-probe-go/cassandra"
	"github.com/andream16/price-probe-go/configuration"
)

func TestCassandraInitialization(t *testing.T) {
	conf := configuration.InitConfiguration()
	_, err := cassandra.InitCassandraClient(conf)
	if err != nil {
		t.Fatalf("Unable to connect to Cassandra!")
	}
}
