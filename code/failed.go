package code

/*
失败code 200000 开头
(200000,201000] -> 基本错误
(201000,202000] -> 认证错误
(202000,203000] -> 注册错误
(203000,204000] -> 登录错误

*/
const (
	BaseFailedCode      = 200000 // 基数失败code
	Failed              = 200001 // 失败
	ParamsError         = 200002 // 参数错误
	DbError             = 200003 // 数据库错误
	IdError             = 200004 // id错误
	AuthFailed          = 201001 // 认证失败
	AuthInvalid         = 201002 // 认证失效
	NoPermission        = 201003 // 没有权限
	UsernameIsExisted   = 202001 // 用户名已经存在
	NameIsExisted       = 202002 // 名称已经存在
	NotChange           = 202003 // 未改变
	LoginFailed         = 203001 // 登录失败
	ActPWdError         = 203002 // 账号/密码错误
	PermissionIsExisted = 700001 // 权限已经存在
	RouterISNotExisted  = 700101 // 路由不存在
)
