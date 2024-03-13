package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	Comment string `json:"comment,omitempty"`

	PostId int `json:"postId,omitempty"`

	Post Post `gorm:"foreignKey:PostId"`
	
}