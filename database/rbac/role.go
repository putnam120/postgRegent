package rbac

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Role struct {
	Name        string   `json:"name"`
	Database    string   `json:"database"`
	Permissions []string `json:"permissions"`
}

func DisablePublic(pool *pgxpool.Pool) error {
	fmt.Println("Removing ALL permissions from the public role.")

	listDatabasesQuery := "SELECT datname FROM pg_database;"
	dbRows, err := pool.Query(context.Background(), listDatabasesQuery)
	if err != nil {
		fmt.Println("Unable to get the list of all databases in the cluster.")
		return err
	}

	var queries []string
	for dbRows.Next() {
		var dbName string
		err := dbRows.Scan(&dbName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Found database", dbName)
		queries = append(queries, fmt.Sprintf("REVOKE ALL ON DATABASE %s FROM public;", dbName))
	}
	queries = append(queries, "REVOKE ALL ON SCHEMA public FROM public;")

	disableQuery := strings.Join(queries[:], "")
	_, err = pool.Exec(context.Background(), disableQuery)

	if err != nil {
		return err
	}

	return nil
}

func CreateRole(pool *pgxpool.Pool, role *Role) error {
	fmt.Println("Creating role ", role.Name)
	var queries []string
	queries = append(queries, fmt.Sprintf("CREATE ROLE role_%s INHERIT;", role.Name))
	for _, permission := range role.Permissions {
		query := fmt.Sprintf("GRANT permission_%s TO role_%s;", permission, role.Name)
		queries = append(queries, query)
	}

	createQuery := strings.Join(queries[:], "")
	_, err := pool.Exec(context.Background(), createQuery)
	if err != nil {
		return err
	}

	fmt.Println("Created role ", role.Name)
	return nil
}
