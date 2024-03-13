package models

import "gorm.io/gorm"

type Following struct {
	gorm.Model
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	UserName string `json:"userName,omitempty"`
}
