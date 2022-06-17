package common

const (

	// headers
	Authorization = "Authorization" // 认证header token

	//schema
	HttpPrefix  = "http://"
	HttpsPrefix = "https://"

	// db signal
	PermissionFindChildren       = "Shard_PermissionFindChildren"
	PermissionRoleFindChildrenId = "Shard_PermissionRoleFindChildrenId"
	RolePermissionFindChildrenId = "Shard_RolePermissionFindChildrenId"

	// route  separator
	RouteSeparator = "||"

	// redis key
	RedisRoutesKey = "shard_routes_%v"
	AdminLoginKey  = "shard_admin_login_%v"
	UserLoginKey   = "shard_user_login_%v"

	// backlash
	Backlash = "/"

	// prefix
	BackendPrefix = "backend"
	ApiPrefix     = "api"

	// version
	V1 = "v1"

	// params
	UriId            = ":id"
	UriAuthorization = ":authorization"

	Socket = "socket"

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
	User        = "user"
	Roles       = "roles"
	Check       = "check"
	File        = "file"
	Host        = "host"
	Log         = "log"
	Task        = "task"
	Stop        = "stop"
	Run         = "run"
	Room        = "room"
)
