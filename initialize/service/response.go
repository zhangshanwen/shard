package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/node"
)

type Res struct {
	StatusCode int         `json:"-"`
	ResCode    int         `json:"code"`
	ReqId      int64       `json:"req_id"`
	Data       interface{} `json:"data"`
	Err        error       `json:"-"`
	Msg        string      `json:"msg"`
}

func success(c *gin.Context, r Res) {
	if r.ResCode == 0 {
		r.ResCode = code.BaseSuccessCode
	}
	if r.Msg == "" {
		r.Msg = common.Ok
	}
	c.JSON(http.StatusOK, r)
}

func failed(c *gin.Context, r Res) {
	if r.ResCode == 0 {
		r.ResCode = code.BaseFailedCode
	}
	r.Msg = r.Err.Error()
	c.JSON(http.StatusBadRequest, r)
}

func Json(c *gin.Context, r Res) {
	r.ReqId = node.N.Generate()
	if r.StatusCode == 0 {
		if r.Err == nil {
			success(c, r)
		} else {
			logrus.Error(r.Err)
			failed(c, r)
		}
		return
	}
	c.JSON(r.StatusCode, r)
}
func (r *Res) AuthFailed() {
	logrus.Error(r.Err)
	r.StatusCode = http.StatusUnauthorized
	r.ResCode = code.AuthFailed
	r.Err = common.AuthFailed
}
func (r *Res) NoPermission() {
	logrus.Error(r.Err)
	r.StatusCode = http.StatusForbidden
	r.ResCode = code.NoPermission
	r.Err = common.NoPermission
}

func (r *Res) DBError() {
	logrus.Error(r.Err)
	r.ResCode = code.DBError
	r.StatusCode = http.StatusBadRequest
	r.Err = common.DBError
}

func (r *Res) NotFound() {
	logrus.Error(r.Err)
	r.ResCode = code.NotFound
	r.StatusCode = http.StatusBadRequest
	r.Err = common.NotFound
}

func (r *Res) NotChange() {
	r.ResCode = code.NotChange
	r.StatusCode = http.StatusBadRequest
	r.Err = common.NotChange

}
func (r *Res) ParamsError() {
	logrus.Error(r.Err)
	r.ResCode = code.ParamsError
	r.StatusCode = http.StatusBadRequest
	r.Err = common.ParamsError
}

func (r *Res) SetPasswordError() {
	logrus.Error(r.Err)
	r.ResCode = code.SetPasswordError
	r.StatusCode = http.StatusBadRequest
	r.Err = common.SetPasswordError
}
func (r *Res) UsernameIsExisted() {
	r.ResCode = code.UsernameIsExisted
	r.StatusCode = http.StatusBadRequest
	r.Err = common.UsernameExisted
}
func (r *Res) NameIsExisted() {
	r.ResCode = code.NameIsExisted
	r.StatusCode = http.StatusBadRequest
	r.Err = common.NameIsExisted
}

func (r *Res) CopierError() {
	logrus.Error(r.Err)
	r.ResCode = code.CopierError
	r.StatusCode = http.StatusBadRequest
	r.Err = common.CopierError
}

func (r *Res) ActPWdError() {
	logrus.Error(r.Err)
	r.ResCode = code.CopierError
	r.StatusCode = http.StatusUnauthorized
	r.Err = common.ActPWdError
}
func (r *Res) LoginFailed() {
	logrus.Error(r.Err)
	r.ResCode = code.LoginFailed
	r.StatusCode = http.StatusUnauthorized
	r.Err = common.LoginFailed
}
func (r *Res) UploadFileFailed() {
	logrus.Error(r.Err)
	r.ResCode = code.UploadFileFailed
	r.StatusCode = http.StatusUnauthorized
	r.Err = common.UploadFileFailed
}
func (r *Res) NotOwner() {
	r.ResCode = code.NotOwner
	r.StatusCode = http.StatusUnauthorized
	r.Err = common.NotOwner
}

func (r *Res) SystemError() {
	logrus.Error(r.Err)
	r.ResCode = code.SystemError
	r.StatusCode = http.StatusBadRequest
	r.Err = common.SystemError
}

func (r *Res) RoomExisted() {
	r.ResCode = code.RoomExisted
	r.StatusCode = http.StatusBadRequest
	r.Err = common.RoomExisted
}

func (r *Res) TaskVerifyError() {
	logrus.Error(r.Err)
	r.ResCode = code.TaskVerifyError
	r.StatusCode = http.StatusBadRequest
}

func (r *Res) TaskAddFailed() {
	logrus.Error(r.Err)
	r.ResCode = code.TaskAddFailed
	r.StatusCode = http.StatusBadRequest
	r.Err = common.TaskAddFailed
}

func (r *Res) TaskRunFailed() {
	logrus.Error(r.Err)
	r.ResCode = code.TaskRunFailed
	r.StatusCode = http.StatusBadRequest
	r.Err = common.TaskRunFailed
}

func (r *Res) TaskStopFailed() {
	logrus.Error(r.Err)
	r.ResCode = code.TaskStopFailed
	r.StatusCode = http.StatusBadRequest
	r.Err = common.TaskStopFailed
}
func (r *Res) TaskIsRunning() {
	r.ResCode = code.TaskIsRunning
	r.StatusCode = http.StatusBadRequest
	r.Err = common.TaskIsRunning

}
func (r *Res) TaskIsNotRunning() {
	r.ResCode = code.TaskIsNotRunning
	r.StatusCode = http.StatusBadRequest
	r.Err = common.TaskIsNotRunning

}
func (r *Res) AmountLtZero() {
	r.ResCode = code.AmountLtZero
	r.StatusCode = http.StatusBadRequest
	r.Err = common.AmountLtZero
}
func (r *Res) BalanceLess() {
	r.ResCode = code.BalanceLess
	r.StatusCode = http.StatusBadRequest
	r.Err = common.BalanceLess
}
