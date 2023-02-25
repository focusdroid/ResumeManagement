package models

import "gorm.io/gorm"

/**
 * @name:"focusdroid"
 * @description:黑名单数据库表
 * @author:1.0
 * @time：2023-02-25 16:35:53
**/

type BlackList struct {
	gorm.Model
	UserName string
	Email    string
	IP       string
	UUID     string
	Phone    string
}

func (backlist BlackList) TableName() string {
	return "blacklist"
}
