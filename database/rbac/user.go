package rbac

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Name     string
	Password string
	Roles    []string
}

func CreateUser(pool *pgxpool.Pool, user *User) error {
	fmt.Println("Creating user: ", user.Name)
	var queries []string
	queries = append(queries,
		fmt.Sprintf(
			"CREATE ROLE user_%s WITH INHERIT LOGIN ENCRYPTED PASSWORD %s;",
			user.Name, user.Password,
		),
	)

	for _, role := range user.Roles {
		query := fmt.Sprintf("GRANT role_%s TO user_%s;", role, user.Name)
		queries = append(queries, query)
	}

	createQuery := strings.Join(queries[:], "")

	if _, err := pool.Exec(context.Background(), createQuery); err != nil {
		return err
	}

	return nil
}
