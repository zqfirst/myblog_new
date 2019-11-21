package models

import (
	"time"
)

type MgUser struct {
	UserId     int       `xorm:"not null pk autoincr INT(11)"`
	Username   string    `xorm:"not null comment('用户名') VARCHAR(20)"`
	Password   string    `xorm:"not null default '' VARCHAR(60)"`
	Salt       string    `xorm:"not null default '' VARCHAR(64)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}


