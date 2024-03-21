package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message  string `json:"message" gorm:"not null;"`
	Photo_Id uint   `json:"photo_id" gorm:"not null;"`
	User_Id  uint   `json:"user_id" gorm:"not null;"`

	Photo Photo `gorm:"foreignKey:photo_id;references:id;"`
	User  User  `gorm:"foreignKey:user_id;references:id;"`
}
