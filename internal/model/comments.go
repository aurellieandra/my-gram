package model

import "time"

type Comment struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement"`
	Message   string     `json:"message" gorm:"not null;" validate:"required"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`

	Photo_Id uint `json:"photo_id" gorm:"column:photo_id;not null;foreignKey:Photo_Id;references:photos(ID)"`
	User_Id  uint `json:"user_id" gorm:"column:user_id;not null;foreignKey:User_Id;references:users(ID)"`
}
