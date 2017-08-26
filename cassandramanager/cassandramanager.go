package cassandramanager

import (
	"errors"

	"github.com/gocql/gocql"
)

func InitCassandraClient() (gocql.Session, error) {

	cluster := gocql.NewCluster("54.194.199.230:9042")
	cluster.Keyspace = "price_probe"
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
