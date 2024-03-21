package model

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title  		string  `json:"title"`
	Url	  		string  `json:"url"`
	Caption   	string	`json:"caption"`
	User_Id     uint 	`json:"user_id" gorm:"not null;foreignKey:user_id;references:id;"`
}