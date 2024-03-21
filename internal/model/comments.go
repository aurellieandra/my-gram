package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message  	string   `json:"message" gorm:"not null;"`
	Photo_Id	uint   	 `json:"photo_id" gorm:"not null;foreignKey:photo_id;references:id;"`
	User_Id     uint 	 `json:"user_id" gorm:"not null;foreignKey:user_id;references:id;"`
}