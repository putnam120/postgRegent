package rbac

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-set"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action string

type actions struct {
	CONNECT Action
	READ    Action
}

type Permission struct {
	Name     string   `json:"name"`
	Database string   `json:"database"`
	Schemas  []string `json:"schemas"`
	Actions  []Action `json:"actions"`
	Tables   []string `json:"tables"`
}

var Actions = actions{
	CONNECT: "CONNECT",
	READ:    "READ",
}

func CreatePermission(pool *pgxpool.Pool, permission *Permission) error {
	var queries []string
	queries = append(queries, fmt.Sprintf("CREATE ROLE permission_%s;", permission.Name))

	// Check that CONNECT is in the set of actions.
	actionsSet := set.From[Action](permission.Actions)
	if actionsSet.Contains(Actions.CONNECT) {
		fmt.Println("Action set contains CONNECT")
		queries = append(queries,
			fmt.Sprintf(
				"GRANT CONNECT ON DATABASE %s TO permission_%s;",
				permission.Database, permission.Name,
			),
		)
	}

	// Add ability to read data in the given (database, schema, table) tuples.
	if actionsSet.Contains(Actions.READ) {
		fmt.Println("Aciton set contains READ")
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
	}

	createQuery := strings.Join(queries[:], "")
	_, err := pool.Exec(context.Background(), createQuery)

	if err != nil {
		return err
	}

	fmt.Println("Created permission ", permission.Name)
	return nil
}
