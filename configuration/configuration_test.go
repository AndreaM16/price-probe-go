package configuration

import (
	"reflect"
	"testing"
)

func TestConfiguration(t *testing.T) {
	config := InitConfiguration()
	serverHost := config.Server.Host
	serverHostType := reflect.TypeOf(serverHost).Kind()
	if serverHostType != reflect.String {
		t.Fatalf("Bad entry for config.Server.Host. Expecting reflect.String, got: " + reflect.TypeOf(serverHost).String())
	}
	if len(serverHost) == 0 {
		t.Fatalf("Bad entry for config.Server.Host. Got empty entry.")
	}
	serverPort := config.Server.Port
	serverPortType := reflect.TypeOf(serverPort).Kind()
	if serverPortType != reflect.Int {
		t.Fatalf("Bad entry for config.Server.Host. Expecting reflect.Int, got: " + reflect.TypeOf(serverPort).String())
	}
	cassandraHost := config.Cassandra.Host
	cassandraHostType := reflect.TypeOf(cassandraHost).Kind()
	if cassandraHostType != reflect.String {
		t.Fatalf("Bad entry for config.Cassandra.Host. Expecting reflect.String, got: " + reflect.TypeOf(cassandraHost).String())
	}
	if len(cassandraHost) == 0 {
		t.Fatalf("Bad entry for config.Cassandra.Host. Got empty entry.")
	}
	cassandraPort := reflect.TypeOf(config.Cassandra.Port).Kind()
	if cassandraPort != reflect.Int {
		t.Fatalf("Bad entry for config.Cassandra.Port. Expecting reflect.Int, got: " + reflect.TypeOf(config.Cassandra.Port).String())
	}
}
