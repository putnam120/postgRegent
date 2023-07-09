package rbac

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Role struct {
	Name        string
	Database    string
	Permissions []string
}

func DisablePublic(pool *pgxpool.Pool, dbname string) error {
	fmt.Println("Removing ALL permissions from the public role.")

	queries := []string{
		"REVOKE ALL ON SCHEMA public FROM PUBLIC;",
		fmt.Sprintf("REVOKE ALL ON DATABASE %s FROM PUBLIC;", dbname),
	}

	disableQuery := strings.Join(queries[:], "")
	_, err := pool.Exec(context.Background(), disableQuery)

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
		query := fmt.Sprintf("GRANT permission_%s TO Irole_%s;", permission, role.Name)
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
