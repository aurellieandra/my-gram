package model

import (
	"time"
)

type User struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement"`
	Username  string     `json:"username" gorm:"column:username;not null;unique;" validate:"unique,required"`
	Email     string     `json:"email" gorm:"column:email;not null;unique;" validate:"email,unique,required"`
	Password  string     `json:"password" gorm:"column:password;not null;" validate:"min=6,required"`
	Dob       time.Time  `json:"dob" gorm:"column:dob;" validate:"required"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
}

type UserResponse struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey;unique;autoIncrement"`
	Username  string     `json:"username" gorm:"column:username;not null;unique;"`
	Email     string     `json:"email" gorm:"column:email;not null;unique;"`
	Dob       time.Time  `json:"dob" gorm:"column:dob;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime;"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;"`
}
