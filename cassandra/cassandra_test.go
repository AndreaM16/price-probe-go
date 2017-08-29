package cassandra

import (
	"testing"

	"github.com/andream16/price-probe-go/configuration"
)

func TestCassandra(t *testing.T) {
	conf := configuration.InitConfiguration()
	session, err := InitCassandraClient(&conf)
	if err != nil {
		t.Fatalf("Unable to connect to Cassandra!")
	}
	DisconnectFromCassandra(session)
}
