package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);unique;not null;column:username"`
	Email    string `gorm:"type:varchar(255);unique;not null;column:email"`
	Password string `gorm:"type:varchar(255);not null;column:password"`
	Role     string `gorm:"type:varchar(255);not null;default:'BASIC';not null;column:role"`
}
