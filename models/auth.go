package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	User          string `gorm:"type:varchar(255)" json:"user"`
	Email         string `gorm:"type:varchar(255)" json:"email"`
	Phone         int    `gorm:"type:varchar(11)" json:"phone"`
	AuthType      string `gorm:"type:varchar(255)" json:"auth_type"`
	AuthNumber    int    `gorm:"type:varchar(255)" json:"auth_number"`
	AuthInterface string `gorm:"type:varchar(255)" json:"auth_interface"`
}

func (auth Auth) TableName() string {
	return "auth"
}
