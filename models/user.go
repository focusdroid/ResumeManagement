package models

import (
	"gorm.io/gorm"
	"time"
)

/*
*
用户表
*/
type User struct {
	gorm.Model
	UserId     string     `gorm:"column:user_id;type:varchar(255);" db:"column:user_id" form:"user_id" json:"user_id" comment:"用户ID"`
	Name       string     `gorm:"column:name;type:varchar(50);not null;" db:"column:name" form:"name" json:"name" comment:"姓名"`
	Phone      string     `gorm:"column:phone;type:varchar(11);" db:"column:phone" form:"phone" json:"phone" comment:"手机号"`
	Email      string     `gorm:"column:email;type:varchar(255);" db:"column:email" json:"email" form:"email" comment:"邮箱"`
	Password   string     `gorm:"column:pass_word;type:varchar(100);" db:"column:pass_word" json:"pass_word" form:"pass_word" comment:"密码"`
	Gender     string     `gorm:"column:gender;type:varchar(10);" db:"column:gender" form:"gender" json:"gender" comment:"性别男M女F"`
	AvatarUrl  string     `gorm:"column:avatar_url;type:varchar(255)" db:"column:avatar_url" form:"avatar_url" json:"avatar_url" comment:"图像"`
	Country    string     `gorm:"column:country;type:varchar(100);" db:"column:country" form:"country" json:"country" comment:"国家"`
	Province   string     `gorm:"column:province;type:varchar(100);" db:"column:province" form:"province" json:"province" comment:"省份"`
	City       string     `gorm:"column:city;type:varchar(100);" db:"column:city" form:"city" json:"city" comment:"城市"`
	NickName   string     `gorm:"column:nick_name;type:varchar(100)" db:"column:nick_name" form:"nick_name" json:"nick_name" comment:"昵称"`
	UserStatus string     `gorm:"column:user_status;type:varchar(100)" db:"column:user_status" form:"user_status" json:"user_status" comment:"用户状态0禁止登录1正常2限制查看人员简历"`
	IsAdmin    bool       `gorm:"column:is_admin;type:tinyint;default:0;" db:"column:is_admin" form:"is_admin" json:"is_admin" comment:"是否是管理员0否1是"`
	IsDelete   bool       `gorm:"column:is_delete;type:tinyint;default:0;" db:"column:is_delete" form:"is_delete" json:"is_delete" comment:"删除用户true"`
	LineTime   *time.Time `gorm:"column:line_time" db:"column:line_time" form:"line_time" json:"line_time" comment:"在线时间"`
}

type UserField struct {
	ID         uint       `json:"id"`
	UserId     string     `json:"user_id"`
	Name       string     `json:"name"`
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	Gender     string     `json:"gender"`
	AvatarUrl  string     `json:"avatar_url"`
	Country    string     `json:"country"`
	Province   string     `json:"province"`
	City       string     `json:"city"`
	NickName   string     `json:"nick_name"`
	UserStatus string     `json:"user_status"`
	IsAdmin    bool       `json:"is_admin"`
	LineTime   *time.Time `json:"line_time"`
}

func (user *User) TableName() string {
	return "user"
}
