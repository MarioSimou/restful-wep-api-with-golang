package models

import (
	"time"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	Id int	 `gorm:"auto_increment;primary key"`
	Username string `gorm:"type:varchar(255);unique;not null;column:username"`
	Email    string `gorm:"type:varchar(255);unique;not null;column:email"`
	Password string `gorm:"type:varchar(255);not null;column:password"`
	Role     string `gorm:"type:varchar(255);not null;default:'BASIC';not null;column:role"`
	CreatedAt time.Time `gorm:"not null;column:created_at"`
	UpdatedAt time.Time `gorm:"not null;column:updated_at"`
}
