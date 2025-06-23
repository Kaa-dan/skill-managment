package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	FullName string
	Email    string
}
