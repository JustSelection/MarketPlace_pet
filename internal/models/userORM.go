package models

import "gorm.io/gorm"

type User struct {
	ID          string         `json:"user_id" gorm:"column:user_id;primaryKey"`
	Name        string         `json:"name,omitempty"`
	Information string         `json:"information,omitempty"`
	Email       string         `json:"email"`
	Password    string         `json:"password,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserUpdate struct {
	Name        string         `json:"name,omitempty"`
	Information string         `json:"information,omitempty"`
	Email       string         `json:"email,omitempty"`
	Password    string         `json:"password,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
