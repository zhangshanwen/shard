package user

import (
	"errors"
	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func BalanceAdjust(c *service.AdminContext) (resp service.Res) {
	pId := param.UriId{}
	if resp.Err = c.BindUri(&pId); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	p := param.AdjustBalance{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{}
	g := db.G.Begin()
	defer func() {
		if resp.Err != nil {
			g.Rollback()
		} else {
			g.Commit()
		}
	}()
	if resp.Err = g.Preload("Wallet").First(&user, pId.Id).Error; resp.Err != nil {
		return
	}
	var recordType model.WalletRecordType
	if p.Amount < 0 {
		recordType = model.WalletRecordTypeBackendDecrease
	} else {
		recordType = model.WalletRecordTypeBackendIncrease
	}
	record := model.WalletRecord{Amount: p.Amount, RecordType: recordType}
	if user.Wallet == nil {
		if p.Amount < 0 {
			resp.ResCode = code.AmountLtZero
			resp.Err = errors.New("AmountLtZero")
			return
		}
		user.Wallet = &model.Wallet{Uid: user.Id, Balance: p.Amount}
		if resp.Err = g.Create(&user.Wallet).Error; resp.Err != nil {
			return
		}
		record.WalletId = user.Wallet.Id
	} else {
		if user.Wallet.Balance+p.Amount < 0 {
			resp.ResCode = code.BalanceLess
			resp.Err = errors.New("BalanceLess")
			return
		}
		if resp.Err = g.Model(&user.Wallet).Update("balance", gorm.Expr("balance + ? ", p.Amount)).Error; resp.Err != nil {
			return
		}
	}
	if resp.Err = g.Create(&record).Error; resp.Err != nil {
		return
	}
	resp.Data = response.PasswordResponse{Password: conf.C.ResetPassword}
	return
}
