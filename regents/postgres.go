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
	Oracle   oracles.Oracle
	connPool *pgxpool.Pool
}

func (r *PostgresRegent) Init() error {
	if err := r.connect(); err != nil {
		return err
	}
	if err := rbac.DisablePublic(r.connPool); err != nil {
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

// Add a new permission to the database.
//
// When attempting to add a permission, an attempt is first made to add it to
// the oracle. If this succeeds we then proceed to add it to the database.
func (r *PostgresRegent) CreatePermission(permission rbac.Permission) error {
	err := r.Oracle.OfferPermission(&permission)
	if err != nil {
		fmt.Printf("Unable to add permission: %v to the oracle.", permission)
		return err
	}
	err = rbac.CreatePermission(r.connPool, &permission)
	if err != nil {
		fmt.Printf("Unable to add permission: %v to the database.", permission)
		return err
	}
	return nil
}
