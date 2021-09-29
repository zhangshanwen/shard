package common

const (
	// headers
	Authorization = "Authorization" // 认证header token

	// db signal
	PermissionFindChildren       = "Shard_PermissionFindChildren"
	PermissionRoleFindChildrenId = "Shard_PermissionRoleFindChildrenId"
	RolePermissionFindChildrenId = "Shard_RolePermissionFindChildrenId"

	// route  separator
	RouteSeparator = "||"

	// redis key
	RedisRoutesKey = "shard_routes_%v"

	// backlash
	Backlash = "/"

	// prefix
	BackendPrefix = "backend"
	ApiPrefix     = "api"

	// version
	V1 = "v1"

	// params
	UriId = ":id"

	// router
	UriEmpty    = ""
	UriLogin    = "login"
	Routes      = "routes"
	UriAvatar   = "avatar"
	Admins      = "admins"
	Password    = "password"
	Balance     = "balance"
	Reset       = "reset"
	Adjust      = "adjust"
	Change      = "change"
	Permissions = "permissions"
	Oss         = "oss"
	Token       = "token"
	Users       = "users"
	Roles       = "roles"
	Check       = "check"
)
