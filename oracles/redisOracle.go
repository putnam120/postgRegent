package oracles

import (
	"fmt"
	"strings"

	"github.com/RediSearch/redisearch-go/v2/redisearch"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
	"github.com/putnam120/postgRegent/database/rbac"
)

type redisOracle struct {
	host        string
	port        int
	username    string
	password    string
	initialized bool
	connPool    *redis.Pool
	searcher    *redisearch.Client
}

func GetRedisOracle(host string, port int, username string, password string) *redisOracle {
	oracle := &redisOracle{host: host, port: port, username: username, password: password}
	oracle.initialized = false
	return oracle
}

func (ro *redisOracle) Init() error {
	if ro.initialized {
		return nil
	}
	// Connect to Redis instance.
	hostPath := fmt.Sprintf("%s:%d", ro.host, ro.port)
	// Create connection pool to Redis.
	ro.connPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", hostPath)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	// Create RediSearch index.
	fmt.Println("Creating RediSearch index.")
	rsearch := redisearch.NewClientFromPool(ro.connPool, "rbacIdx")
	idxDef := &redisearch.IndexDefinition{IndexOn: "JSON", Prefix: []string{"permission:", "role:", "user:"}}
	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("$.name", redisearch.TextFieldOptions{As: "name"})).
		AddField(redisearch.NewTagFieldOptions("$.permissions", redisearch.TagFieldOptions{As: "permissions"})).
		AddField(redisearch.NewTagFieldOptions("$.roles", redisearch.TagFieldOptions{As: "roles"}))

	err := rsearch.CreateIndexWithIndexDefinition(schema, idxDef)
	if err != nil {
		fmt.Println("Possible error while creating RedisOracle: ", err)
		message := err.Error()
		if !strings.Contains(message, "Index already exists") {
			return err
		}
	}
	ro.searcher = rsearch
	ro.initialized = true
	fmt.Println("Successfully initialized RedisOracle.")
	return nil
}

// Return if there exists a document with the name 'prefix_name'.
func (ro *redisOracle) searchForName(prefix string, name string) bool {
	query := fmt.Sprintf("'@name:%s_%s'", prefix, name)
	_, num, err := ro.searcher.Search(redisearch.NewQuery(query))
	if err != nil {
		fmt.Println("Error while searching for ", prefix, "(", name, "): ", err)
		return false
	}
	return num > 0
}

// Obtain the name of documents that have the requested tagName.
func (ro *redisOracle) searchForTag(fieldName string, tagName string) ([]string, error) {
	query := fmt.Sprintf("'@%s:{%s}'", fieldName, tagName)
	docs, _, err := ro.searcher.Search(redisearch.NewQuery(query).AddReturnFields("name"))
	if err != nil {
		return nil, err
	}
	var names []string
	for _, doc := range docs {
		names = append(names, doc.Properties["name"].(string))
	}
	return names, nil
}

func (ro *redisOracle) PermissionExists(name string) bool {
	return ro.searchForName("permission", name)
}

func (ro *redisOracle) RoleExists(name string) bool {
	return ro.searchForName("role", name)
}

func (ro *redisOracle) UserExists(name string) bool {
	return ro.searchForName("user", name)
}

func (ro *redisOracle) RolesWithPermission(name string) ([]string, error) {
	return ro.searchForTag("permissions", name)
}

func (ro *redisOracle) UsersWithRole(name string) ([]string, error) {
	return ro.searchForTag("roles", name)
}

// Give the Oracle informaiton about a permission.
//
// If the permission already exists no action is taken. Otherwise,
// the data is added to the Redis databse. Any errors that are encountered
// are returned.
func (ro *redisOracle) OfferPermission(permission *rbac.Permission) error {
	updatedPermission := *permission
	updatedPermission.Name = "permission_" + permission.Name
	if ro.PermissionExists(permission.Name) {
		fmt.Println("Permission ", permission.Name, " already exists.")
		return nil
	}
	// Use RedisJSON to insert the document.
	conn := ro.connPool.Get()
	defer conn.Close()
	rjson := rejson.NewReJSONHandler()
	rjson.SetRedigoClient(conn)
	key := "permission:" + updatedPermission.Name
	_, err := rjson.JSONSet(key, ".", updatedPermission)
	return err
}

// Give the Oracle informaiton about a role.
//
// If the role already exists no action is taken. Otherwise,
// the data is added to the Redis databse. Any errors that are encountered
// are returned.
func (ro *redisOracle) OfferRole(role *rbac.Role) error {
	updatedRole := *role
	updatedRole.Name = "role_" + role.Name

	if ro.RoleExists(role.Name) {
		fmt.Println("Role ", role.Name, " already exists.")
		return nil
	}
	// Ensure that all of the listed permissions exist.
	permissions := []string{}
	for _, permission := range role.Permissions {
		permissions = append(permissions, "permission_"+permission)
	}
	query := fmt.Sprintf("'@name:(%s)'", strings.Join(permissions, "|"))
	_, count, err := ro.searcher.Search(redisearch.NewQuery(query))
	if err != nil {
		return err
	}
	if count != len(role.Permissions) {
		return fmt.Errorf("Not all of the provided permissions exist.")
	}

	// Use RedisJSON to insert the document.
	conn := ro.connPool.Get()
	defer conn.Close()
	rjson := rejson.NewReJSONHandler()
	rjson.SetRedigoClient(conn)
	key := "role:" + updatedRole.Name
	_, err = rjson.JSONSet(key, ".", updatedRole)
	return err
}

// Give the Oracle informaiton about a user.
//
// If the user already exists no action is taken. Otherwise,
// the data is added to the Redis databse. Any errors that are encountered
// are returned.
func (ro *redisOracle) OfferUser(user *rbac.User) error {
	updatedUser := *user
	updatedUser.Name = "user_" + user.Name
	if ro.UserExists(user.Name) {
		fmt.Println("User ", user.Name, " already exists.")
		return nil
	}
	// Ensure that all of the listed roles exist.
	roles := []string{}
	for _, role := range user.Roles {
		roles = append(roles, "role_"+role)
	}
	query := fmt.Sprintf("'@name:(%s)'", strings.Join(roles, "|"))
	_, count, err := ro.searcher.Search(redisearch.NewQuery(query))
	if err != nil {
		return err
	}
	if count != len(user.Roles) {
		return fmt.Errorf("Not all of the provided roles exist.")
	}
	// Use RedisJSON to insert the document.
	conn := ro.connPool.Get()
	defer conn.Close()
	rjson := rejson.NewReJSONHandler()
	rjson.SetRedigoClient(conn)
	key := "user:" + updatedUser.Name
	_, err = rjson.JSONSet(key, ".", updatedUser)
	return err
}
