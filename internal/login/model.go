package main

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	UserName string
	Phone    *string `gorm:"not null"`
	Email    string  `gorm:"type:varchar(100);unique_index"`
	ID       int     `gorm:"AUTO_INCREMENT"`
}
