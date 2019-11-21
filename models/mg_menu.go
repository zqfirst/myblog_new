package models

import (
	"time"
)

type MgMenu struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	Name       string    `xorm:"not null comment('菜单名') VARCHAR(20)"`
	Icon       string    `xorm:"not null default '' comment('icon 类') VARCHAR(60)"`
	Pid        int       `xorm:"not null default 0 comment('父级id') INT(11)"`
	Sort       int       `xorm:"not null default 0 INT(11)"`
	Url        string    `xorm:"not null default '' VARCHAR(255)"`
	Status     int       `xorm:"not null default 0 comment('1启用2禁用') TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
