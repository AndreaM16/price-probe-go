package cassandra

import (
	"errors"

	"github.com/andream16/price-probe-go/configuration"
	"github.com/gocql/gocql"
)

// InitCassandraClient takes a configuration files
// return cassandra singleton to gocql.Session
func InitCassandraClient(conf *configuration.Configuration) (gocql.Session, error) {
	cluster := gocql.NewCluster(conf.Cassandra.Host)
	cluster.Keyspace = conf.Cassandra.Keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return *session, errors.New("Cannot Initialize Cassandra Cluster. Got: " + err.Error())
	}
	return *session, err
}

// DisconnectFromCassandra takes the current gocql.Session
// disconnects from it
func DisconnectFromCassandra(session gocql.Session) {
	session.Close()
	return
}
