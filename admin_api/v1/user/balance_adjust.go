package user

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func BalanceAdjust(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.AdjustBalance{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		u  model.User
		tx = c.Tx
	)
	if r.Err = tx.Preload("Wallet").First(&u, pId.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	var recordType model.WalletRecordType
	if p.Amount < 0 {
		recordType = model.WalletRecordTypeBackendDecrease
	} else {
		recordType = model.WalletRecordTypeBackendIncrease
	}
	record := model.WalletRecord{Amount: p.Amount, RecordType: recordType}
	if u.Wallet == nil {
		if p.Amount < 0 {
			r.AmountLtZero()
			return
		}
		u.Wallet = &model.Wallet{Uid: u.Id, Balance: p.Amount}
		if r.Err = tx.Create(&u.Wallet).Error; r.Err != nil {
			r.DBError()
			return
		}
		record.WalletId = u.Wallet.Id
	} else {
		if u.Wallet.Balance+p.Amount < 0 {
			r.BalanceLess()
			return
		}
		if r.Err = tx.Model(&u.Wallet).Update("balance", gorm.Expr("balance + ? ", p.Amount)).Error; r.Err != nil {
			r.DBError()
			return
		}
	}
	if r.Err = tx.Create(&record).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("修改用户(id:%v,username:%v)余额 %v ,修改后余额 %v",
		u.Id, u.Username, p.Amount, u.Wallet.Balance), model.OperateLogTypeUpdate)
	return
}
