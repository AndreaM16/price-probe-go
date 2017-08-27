package main

import (
	"testing"

	"../cassandramanager"
)

func TestCassandraInitialization(t *testing.T) {
	_, err := cassandramanager.InitCassandraClient()
	if err != nil {
		t.Fatalf("Unable to connect to Cassandra!")
	}
}
