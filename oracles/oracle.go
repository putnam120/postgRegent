package oracles

import "github.com/putnam120/postgRegent/database/rbac"

/*
Interface for a component that the Regent uses
to obtain informaiton about the roles in the Postgres
database.
*/
type Oracle interface {
	Init() error

	PermissionExists(string) bool
	RoleExists(string) bool
	UserExists(string) bool

	RolesWithPermission(string) ([]string, error)
	UsersWithRole(string) ([]string, error)

	OfferPermission(*rbac.Permission) error
	OfferRole(*rbac.Role) error
	OfferUser(*rbac.User) error
}
