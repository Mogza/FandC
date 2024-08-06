package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email" gorm:"unique"`
}
