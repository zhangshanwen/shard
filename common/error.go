package common

import "errors"

var (
	AuthFailed       = "AuthFailed"
	TaskIsRunning    = "TaskIsRunning"
	TaskIsNotRunning = "TaskIsNotRunning"
	UsernameExisted  = "UsernameExisted"
	NameIsExisted    = "NameIsExisted"
	ActPWdError      = "ActPWdError"
	NotOwner         = "NotOwner"
	ParamsError      = "ParamsError"
	DBError          = "DBError"
	NotFound         = "NotFound"
	SetPasswordError = "SetPasswordError"
	CopierError      = "CopierError"
	LoginFailed      = "LoginFailed"
	UploadFileFailed = "UploadFileFailed"
	SystemError      = "SystemError"
	RoomExisted      = "RoomExisted"
	NotChange        = "NotChange"
	TaskAddFailed    = "TaskAddFailed"
	TaskStopFailed   = "TaskStopFailed"
	AmountLtZero     = "AmountLtZero"
	BalanceLess      = "BalanceLess"
	NoPermission     = "NoPermission"
	TaskRunFailed    = "TaskRunFailed"
)

var (
	IdErr                                       = errors.New("IdErr")
	TaskRunErr                                  = errors.New("TaskRunErr")
	RequestErr                                  = errors.New("RequestErr")
	NoPermissionErr                             = errors.New("NoPermission")
	ExpiryTimeShouldNotBeLessThanOneDayErr      = errors.New("ExpiryTimeShouldNotBeLessThanOneDay")      // 过期时间不得小于1天
	ExpiryTimeShouldNotBeLessThanCurrentTimeErr = errors.New("ExpiryTimeShouldNotBeLessThanCurrentTime") // 过期时间不得小于当前时间

)
