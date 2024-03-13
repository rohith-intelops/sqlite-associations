package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	Date int `json:"date,omitempty"`

	Message string `json:"message,omitempty"`

	UserId int `json:"userId,omitempty"`

}