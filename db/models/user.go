// db/models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email" gorm:"type:varchar(255);uniqueIndex:idx_email,unique" validate:"required,email"`
	Username  string `json:"username" gorm:"type:varchar(255);uniqueIndex:idx_username,unique" validate:"required,min=3"`
	Password  string `json:"password" validate:"required,min=6"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
}
