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
	EDIT    Action
	ADMIN   Action
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
	EDIT:    "EDIT",
	ADMIN:   "ADMIN",
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
	if actionsSet.Contains(Actions.READ) &&
		!(actionsSet.Contains(Actions.EDIT) || actionsSet.Contains(Actions.ADMIN)) {

		fmt.Println("Aciton set for a READER")
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

	// Add ability to edit (but not delete) data in the given (database, schema, table) tubles.
	if actionsSet.Contains(Actions.EDIT) && !actionsSet.Contains(Actions.ADMIN) {
		fmt.Println("Action set for an EDITOR")
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
	}

	if actionsSet.Contains(Actions.ADMIN) {
		fmt.Println("Action set for an ADMIN")
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

	}

	createQuery := strings.Join(queries[:], "")
	_, err := pool.Exec(context.Background(), createQuery)

	if err != nil {
		return err
	}

	fmt.Println("Created permission ", permission.Name)
	return nil
}
