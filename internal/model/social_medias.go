package model

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name  		string  `json:"name"`
	Url	  		string  `json:"url"`
	User_Id     uint 	`json:"user_id" gorm:"not null;foreignKey:user_id;references:id;"`
}