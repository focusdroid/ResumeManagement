package models

import "gorm.io/gorm"

type Backlog struct {
	gorm.Model
	UserId        string `gorm:"column:user_id;type:varchar(255);" db:"column:user_id" form:"user_id" json:"user_id" comment:"用户ID"`
	BacklogText   string `gorm:"column:backlog_text;type:text;" db:"column:backlog_text" form:"backlog_text" json:"backlog_text" comment:"待办列表文本"`
	BacklogStatus int    `gorm:"column:backlog_status;type:tinyint;not null;default:1" db:"column:backlog_status" form:"backlog_status" json:"backlog_status" comment:"已删除0正常显示1轻度紧急2中度紧急3非常紧急4"`
	BacklogType   int    `gorm:"column:backlog_type;type:tinyint;not null;default:1" db:"column:backlog_type" form:"backlog_type" json:"backlog_type" comment:"正在待办1,已完成2"`
}

type BacklogInterface struct {
	Id            uint   `json:"id"`
	UserId        string `json:"user_id"`
	BacklogText   string `json:"backlog_text"`
	BacklogStatus int    `json:"backlog_status"`
	BacklogType   int    `json:"backlog_type"`
}

func (backlog Backlog) TableName() string {
	return "backlog"
}
