package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username" gorm:"primaryKey;unique;"`
	Email	  string    `json:"email" gorm:"not null;unique;"`
	Password  string    `json:"password" gorm:"not null;"`
	DoB       time.Time `json:"dob"`
}