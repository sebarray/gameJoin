package dto

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	PasswordHash string `json:"password"`
}
