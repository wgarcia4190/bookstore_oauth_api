package cassandra

import (
	"os"

	"github.com/gocql/gocql"
)

const (
	cUsername = "CASSANDRA_SEEDS"
	cPassword = "CASSANDRA_PASSWORD"
)

var (
	cluster *gocql.ClusterConfig

	username = os.Getenv(cUsername)
	password = os.Getenv(cPassword)
)

func init() {
	cluster = gocql.NewCluster("127.0.0.1")

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
