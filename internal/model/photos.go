package model

import "time"

type Photo struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement"`
	Title     string     `json:"title" gorm:"column:title;not null;" validate:"required"`
	Url       string     `json:"url" gorm:"column:url;not null; validate:"required"`
	Caption   string     `json:"caption" gorm:"column:caption;not null;`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`

	User_Id uint64 `json:"user_id" gorm:"column:user_id;not null;foreignKey:User_Id;references:users(ID)"`
}
