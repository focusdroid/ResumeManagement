package models

import "gorm.io/gorm"

type Backlog struct {
	gorm.Model
	UserId        string `gorm:"column:user_id;type:varchar(255);" db:"column:user_id" form:"user_id" json:"user_id" comment:"用户ID"`
	BacklogText   string `gorm:"column:backlog_text;type:text;" db:"column:backlog_text" form:"backlog_text" json:"backlog_text" comment:"待办列表文本"`
	BacklogStatus string `gorm:"column:backlog_status;varchar(20);default:1;" db:"column:backlog_status" form:"backlog_status" json:"backlog_status" comment:"待办状态0已删除1正常显示2轻度紧急3中度紧急4非常紧急"`
}

func (backlog Backlog) TableName() string {
	return "backlog"
}
