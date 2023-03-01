package models

import "gorm.io/gorm"

/**
 * @author: focusdroid
 * @description: 字符串权限表
 * @version: 1.0
 * @time：2023-03-01 14:11:33
**/
type Auth struct {
	gorm.Model
	User          string `gorm:"type:varchar(255)" from:"user" db:"column:user"  json:"user" comment:"用户名称"`
	Email         string `gorm:"type:varchar(255)" from:"email" db:"column:email"  json:"email" comment:"邮箱"`
	Phone         int    `gorm:"type:varchar(11)"  from:"phone" db:"column:phone"  json:"phone" comment:"手机号"`
	AuthType      string `gorm:"type:varchar(255)" from:"auth_type" db:"column:auth_type"  json:"auth_type" comment:"权限类型"`
	AuthNumber    int    `gorm:"type:varchar(255)" from:"auth_number" db:"column:auth_number"  json:"auth_number" comment:"权限等级"`
	AuthInterface string `gorm:"type:varchar(255)" from:"auth_interface" db:"column:auth_interface"  json:"auth_interface" comment:"该用户对应的权限"`
}

func (auth Auth) TableName() string {
	return "auth"
}
