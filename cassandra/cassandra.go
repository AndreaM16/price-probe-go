package cassandra

import (
	"errors"

	"github.com/andream16/price-probe-go/configuration"
	"github.com/gocql/gocql"
)

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

func DisconnectFromCassandra(session gocql.Session) {
	session.Close()
	return
}
