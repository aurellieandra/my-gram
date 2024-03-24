package model

import "time"

type SocialMedia struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement"`
	Name      string     `json:"name" gorm:"column:name;not null;"`
	Url       string     `json:"url" gorm:"column:url;not null;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`

	User_Id uint64 `json:"user_id" gorm:"column:user_id;not null;foreignKey:UserID;references:users(ID)"`
}

func (SocialMedia) TableName() string {
	return "social_medias"
}
