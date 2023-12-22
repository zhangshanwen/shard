package code

/*
失败code 200000 开头
(200000,201000] -> 基本错误
(201000,202000] -> 认证错误
(202000,203000] -> 注册错误
(203000,204000] -> 登录错误
*/
const (
	BaseFailedCode    = 200000 // 基数失败code
	Failed            = 200001 // 失败
	ParamsError       = 200002 // 参数错误
	DBError           = 200003 // 数据库错误
	IdError           = 200004 // id错误
	NotFound          = 200005 // 数据不存在
	CopierError       = 200006 // copier错误
	SystemError       = 200007 // 系统错误
	AuthFailed        = 201001 // 认证失败
	AuthInvalid       = 201002 // 认证失效
	NoPermission      = 201003 // 没有权限
	CreateTokenError  = 201004 // 创建token失败
	UsernameIsExisted = 202001 // 用户名已经存在
	NameIsExisted     = 202002 // 名称已经存在
	AmountLtZero      = 202003 // 数量小于0
	BalanceLess       = 202004 // 余额不足
	NotChange         = 202005 // 未改变
	LoginFailed       = 203001 // 登录失败
	ActPWdError       = 203002 // 账号/密码错误
	SetPasswordError  = 203003 // 设置密码错误
	UploadFileFailed  = 204001 // 上传文件失败
	NotOwner          = 204002 // 不是拥有者
	RoomExisted       = 204003 // 房间已存在
	RoomJoinFailed    = 204004 // 房间加入失败
	PushFailed        = 204005 // 房间加入失败
	TaskVerifyError   = 205001 // 任务校验失败
	TaskAddFailed     = 205002 // 任务添加失败
	TaskStopFailed    = 205003 // 任务停止失败
	TaskRunFailed     = 205004 // 任务执行失败
	TaskIsRunning     = 205005 // 任务正在执行
	TaskIsNotRunning  = 205006 // 任务未运行
	Living            = 206001 // 正在直播

	PermissionIsExisted = 700001 // 权限已经存在
	RouterISNotExisted  = 700101 // 路由不存在
)
