package common

const (

	// headers
	Authorization = "authorization" // 认证header token

	LiveAppName = "live"

	//schema
	HttpPrefix  = "http://"
	HttpsPrefix = "https://"

	// db signal
	PermissionFindChildren       = "Shard_PermissionFindChildren"
	PermissionRoleFindChildrenId = "Shard_PermissionRoleFindChildrenId"
	RolePermissionFindChildrenId = "Shard_RolePermissionFindChildrenId"

	// route  separator
	RouteSeparator = "||"

	MessageSplitSymbol = "[::]"

	//
	TimeFullFormat  = "2006-01-02 15:04:05"
	TimeCrushFormat = "20060102150405"

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

	WechatStorageFilePath = "storage"

	// router
	UriEmpty    = ""
	UriLogin    = "login"
	Logout      = "logout"
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
	Upload      = "upload"
	Users       = "users"
	User        = "user"
	Roles       = "roles"
	Check       = "check"
	File        = "file"
	Host        = "host"
	Live        = "live"
	Meeting     = "meeting"
	Barrage     = "barrage"
	Log         = "log"
	Task        = "task"
	Stop        = "stop"
	Start       = "start"
	Run         = "run"
	Room        = "room"
	Watch       = "watch"
	Empty       = "empty"
	Wechat      = "wechat"
	Qrcode      = "qrcode"
	Push        = "push"
	Join        = "join"
	Offset      = "offset"
	Status      = "status"
	Avatar      = "avatar"
	Rules       = "rules"
	Bots        = "bots"
	Timer       = "timer"
	Reply       = "reply"
	Functions   = "functions"
)
