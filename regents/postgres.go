package regents

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/putnam120/postgRegent/database"
	"github.com/putnam120/postgRegent/database/rbac"
	"github.com/putnam120/postgRegent/oracles"
)

type PostgresRegent struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	connPool *pgxpool.Pool
	Oracle   oracles.Oracle
}

func (r *PostgresRegent) Init() error {
	if err := r.connect(); err != nil {
		return err
	}
	if err := rbac.DisablePublic(r.connPool, r.Dbname); err != nil {
		return err
	}

	fmt.Println("Successfully created regent for host: ", r.Host)
	return nil
}

func (r *PostgresRegent) connect() error {
	if len(r.Dbname) == 0 {
		r.Dbname = r.User
	}
	conn, err := database.Connect(r.Host, r.Port, r.User, r.Password, r.Dbname)

	if err != nil {
		return err
	}

	r.connPool = conn
	return nil
}
