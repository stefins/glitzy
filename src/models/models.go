package models

import (
	"gorm.io/gorm"
)

// User hold the info of a credential
type User struct {
	gorm.Model
	ServiceName string
	Username    string
	Password    string
}

type Main struct {
	gorm.Model
	PasswordHash string
}
