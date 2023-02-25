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
	UserName string `gorm:"column:user_name;type:varchar(255)" db:"column:user_name" form:"user_name" json:"user_name" comment:"用户名称"`
	Email    string `gorm:"column:email;type:varchar(255)" db:"column:email" form:"email" json:"email" comment:"email"`
	IP       string `gorm:"column:ip;type:varchar(20)" db:"column:ip" form:"ip" json:"ip" comment:"ip地址"`
	IPV6     string `gorm:"column:ipv6;type:varchar(100)" db:"column:ipv6" form:"ipv6" json:"ipv6" comment:"ipv6"`
	UUID     string `gorm:"column:uuid;type:varchar(255)" db:"column:uuid" form:"uuid" json:"uuid" comment:"uuid"`
	Phone    string `gorm:"column:phone;type:varchar(11)" db:"column:phone" form:"phone" json:"phone" comment:"电话"`
}

func (backlist BlackList) TableName() string {
	return "blacklist"
}
