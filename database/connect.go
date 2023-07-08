package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var connPoolMap = make(map[string]*pgxpool.Pool)
var poolCreator = make(map[string]*sync.Once)
var connError error

// Obtain a connection pool to a PostgreSQL database.
//
// If a connection pool already exists it is used. Otherwise, a new pool
// is created.
//
// Args:
//   - host (string): Address of the RDBMS hosting the Postgres cluster.
//   - port (int): Port of the desired Postgres instance.
//   - user (string): User connecting to the database.
//   - password (string): User's password.
//   - dbname (string): The name of the database you wish to connect to.
//
// Returns:
//   - (*pgxpool.Pool, error): A connection pool if no errors were encountered. Otherwise an error.
func Connect(host string, port int, user string, password string, dbname string) (*pgxpool.Pool, error) {
	connectionId := fmt.Sprintf("%s-%s", host, dbname)
	creator, ok := poolCreator[connectionId]
	if !ok {
		poolCreator[connectionId] = &sync.Once{}
		creator = poolCreator[connectionId]
	}
	creator.Do(func() {
		dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
		pool, err := pgxpool.New(context.Background(), dbUrl)
		if err != nil {
			_ = fmt.Errorf("Unable to connect to database: %w", err)
			connError = err
		} else {
			fmt.Println("Successfully connected to database ", dbname, " in ", host)
			connPoolMap[connectionId] = pool
		}
	})

	if connError != nil {
		return nil, connError
	}

	return connPoolMap[connectionId], nil
}
