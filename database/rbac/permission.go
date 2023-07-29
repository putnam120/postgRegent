package rbac

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

type Action string

type actions struct {
	CONNECT Action
	READ    Action
	EDIT    Action
	ADMIN   Action
}

type Permission struct {
	Name     string   `json:"name"`
	Database string   `json:"database"`
	Schemas  []string `json:"schemas"`
	Action   Action   `json:"action"`
	Tables   []string `json:"tables"`
}

var Actions = actions{
	CONNECT: "CONNECT",
	READ:    "READ",
	EDIT:    "EDIT",
	ADMIN:   "ADMIN",
}

func CreatePermission(pool *pgxpool.Pool, permission *Permission) error {
	var queries []string
	queries = append(queries, fmt.Sprintf("CREATE ROLE permission_%s;", permission.Name))

	switch permission.Action {
	case Actions.CONNECT:
		fmt.Println("CONNECT permssion")
		queries = append(queries,
			fmt.Sprintf(
				"GRANT CONNECT ON DATABASE %s TO permission_%s;",
				permission.Database, permission.Name,
			),
		)
	case Actions.READ:
		fmt.Println("READ permission")
		queries = append(queries,
			fmt.Sprintf("GRANT SELECT ON %s IN SCHEMA %s TO %s;",
				strings.Join(permission.Tables[:], ","),
				strings.Join(permission.Schemas[:], ","),
				permission.Name,
			),
			fmt.Sprintf("GRANT USAGE ON SCHEMA %s TO %s;",
				strings.Join(permission.Schemas[:], ","),
				permission.Name,
			),
		)
	case Actions.EDIT:
		// Add ability to edit (but not delete) data in the given (database, schema, table) tubles.
		fmt.Println("EDIT permission")
		queries = append(queries,
			fmt.Sprintf("GRANT SELECT,INSERT,UPDATE ON %s IN SCHEMA %s TO %s;",
				strings.Join(permission.Tables[:], ","),
				strings.Join(permission.Schemas[:], ","),
				permission.Name,
			),
			fmt.Sprintf("GRANT USAGE ON SCHEMA %s TO %s;",
				strings.Join(permission.Schemas[:], ","),
				permission.Name,
			),
		)
	case Actions.ADMIN:
		// Same abilities as EDIT but is also able to delete data.
		fmt.Println("ADMIN permission")
		queries = append(queries,
			fmt.Sprintf("GRANT ALL ON %s IN SCHEMA %s TO %s;",
				strings.Join(permission.Tables[:], ","),
				strings.Join(permission.Schemas[:], ","),
				permission.Name,
			),
			fmt.Sprintf("GRANT ALL ON SCHEMA %s TO %s;",
				strings.Join(permission.Schemas[:], ","),
				permission.Name,
			),
		)
	default:
		return fmt.Errorf("Unrecognized permission action: %s", permission.Action)
	}

	createQuery := strings.Join(queries[:], "")
	_, err := pool.Exec(context.Background(), createQuery)

	if err != nil {
		return err
	}

	fmt.Println("Created permission ", permission.Name)
	return nil
}
