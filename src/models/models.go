package models

import (
	"gorm.io/gorm"
)

// User hold the info of a credential
type User struct {
	gorm.Model
	Name     string `gorm:"column:name"`
	Username string
	Password string
}

type Main struct {
	gorm.Model
	PasswordHash string
}
