package service

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
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
	if r.Msg == "" {
		r.Msg = r.Err.Error()
	}
	logrus.Errorf("%s %v \n", r.Msg, r.Err)
	if r.ResCode == 0 {
		r.ResCode = code.BaseFailedCode
	}
	c.JSON(http.StatusBadRequest, r)
}

func Json(c *gin.Context, r Res) {
	r.ReqId = node.N.Generate()
	if r.Data == nil {
		r.Data = struct{}{}
	}
	if r.StatusCode == 0 {
		if r.Msg == "" || r.Err == nil {
			success(c, r)
		} else {
			failed(c, r)
		}
		return
	}
	c.JSON(r.StatusCode, r)
}
func (r *Res) AuthFailed() {
	r.StatusCode = http.StatusUnauthorized
	r.ResCode = code.AuthFailed
	r.Msg = common.AuthFailed
}
func (r *Res) NoPermission() {
	r.StatusCode = http.StatusForbidden
	r.ResCode = code.NoPermission
	r.Msg = common.NoPermission
}

func (r *Res) DBError() {
	r.ResCode = code.DBError
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.DBError
}

func (r *Res) NotFound() {
	r.ResCode = code.NotFound
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.NotFound
}

func (r *Res) NotChange() {
	r.ResCode = code.NotChange
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.NotChange

}
func (r *Res) ParamsError() {
	r.ResCode = code.ParamsError
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.ParamsError
}

func (r *Res) SetPasswordError() {
	r.ResCode = code.SetPasswordError
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.SetPasswordError
}
func (r *Res) UsernameIsExisted() {
	r.ResCode = code.UsernameIsExisted
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.UsernameExisted
}
func (r *Res) NameIsExisted() {
	r.ResCode = code.NameIsExisted
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.NameIsExisted
}

func (r *Res) CopierError() {
	r.ResCode = code.CopierError
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.CopierError
}

func (r *Res) ActPWdError() {
	r.ResCode = code.CopierError
	r.StatusCode = http.StatusUnauthorized
	r.Msg = common.ActPWdError
}
func (r *Res) LoginFailed() {
	r.ResCode = code.LoginFailed
	r.StatusCode = http.StatusUnauthorized
	r.Msg = common.LoginFailed
}
func (r *Res) UploadFileFailed() {
	r.ResCode = code.UploadFileFailed
	r.StatusCode = http.StatusUnauthorized
	r.Msg = common.UploadFileFailed
}
func (r *Res) NotOwner() {
	r.ResCode = code.NotOwner
	r.StatusCode = http.StatusUnauthorized
	r.Msg = common.NotOwner
}

func (r *Res) SystemError() {
	r.ResCode = code.SystemError
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.SystemError
}

func (r *Res) RoomExisted() {
	r.ResCode = code.RoomExisted
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.RoomExisted
}

func (r *Res) TaskVerifyError() {
	r.ResCode = code.TaskVerifyError
	r.StatusCode = http.StatusBadRequest
}

func (r *Res) TaskAddFailed() {
	r.ResCode = code.TaskAddFailed
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.TaskAddFailed
}

func (r *Res) TaskRunFailed() {
	r.ResCode = code.TaskRunFailed
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.TaskRunFailed
}

func (r *Res) TaskStopFailed() {
	r.ResCode = code.TaskStopFailed
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.TaskStopFailed
}
func (r *Res) TaskIsRunning() {
	r.ResCode = code.TaskIsRunning
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.TaskIsRunning

}
func (r *Res) TaskIsNotRunning() {
	r.ResCode = code.TaskIsNotRunning
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.TaskIsNotRunning

}
func (r *Res) AmountLtZero() {
	r.ResCode = code.AmountLtZero
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.AmountLtZero
}
func (r *Res) BalanceLess() {
	r.ResCode = code.BalanceLess
	r.StatusCode = http.StatusBadRequest
	r.Msg = common.BalanceLess
}
func (r *Res) Living() {

}
