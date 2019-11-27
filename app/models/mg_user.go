package models

import (
	"time"
)

type MgUser struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Username   string    `xorm:"not null comment('用户名') VARCHAR(20)"`
	Password   string    `xorm:"not null default '' VARCHAR(60)"`
	Salt       string    `xorm:"not null default '' VARCHAR(64)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}

type AdminOperate struct {
	operate
}

var User AdminOperate

func init() {
	User = AdminOperate{}
}

func (this AdminOperate) GetUser(id string) {
	a := &MgUser{}
	this.QueryOne(id, a)
}
