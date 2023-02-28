package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	User          string
	Email         string
	Phone         int
	AuthType      string
	AuthNumber    int
	AuthInterface string
}

func (auth Auth) TableName() string {
	return "auth"
}
