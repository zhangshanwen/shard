package common

import "errors"

var (
	AuthFailed                               = errors.New("AuthFailed")
	NoPermission                             = errors.New("NoPermission")
	ErrorId                                  = errors.New("ErrorId")
	EditNoChange                             = errors.New("EditNoChange")
	TaskIsRunning                            = errors.New("TaskIsRunning")
	RequestFailed                            = errors.New("RequestFailed")
	TaskIsNotRunning                         = errors.New("TaskIsNotRunning")
	UsernameExisted                          = errors.New("UsernameExisted")
	NameIsExisted                            = errors.New("NameIsExisted")
	ActPWdError                              = errors.New("ActPWdError")
	NotOwner                                 = errors.New("NotOwner")
	ParamsError                              = errors.New("ParamsError")
	DBError                                  = errors.New("DBError")
	NotFound                                 = errors.New("NotFound")
	SetPasswordError                         = errors.New("SetPasswordError")
	CopierError                              = errors.New("CopierError")
	LoginFailed                              = errors.New("LoginFailed")
	UploadFileFailed                         = errors.New("UploadFileFailed")
	SystemError                              = errors.New("SystemError")
	RoomExisted                              = errors.New("RoomExisted")
	NotChange                                = errors.New("NotChange")
	ExpiryTimeShouldNotBeLessThanOneDay      = errors.New("ExpiryTimeShouldNotBeLessThanOneDay")      // 过期时间不得小于1天
	ExpiryTimeShouldNotBeLessThanCurrentTime = errors.New("ExpiryTimeShouldNotBeLessThanCurrentTime") // 过期时间不得小于当前时间
	TaskAddFailed                            = errors.New("TaskAddFailed")
	TaskStopFailed                           = errors.New("TaskStopFailed")
	TaskRunFailed                            = errors.New("TaskRunFailed")
	AmountLtZero                             = errors.New("AmountLtZero")
	BalanceLess                              = errors.New("BalanceLess")
)
