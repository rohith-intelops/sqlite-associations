package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	EmailId string `json:"emailId,omitempty"`

	Name string `json:"name,omitempty"`

	Post []Post 
}
