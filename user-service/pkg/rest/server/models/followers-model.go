package models

import "gorm.io/gorm"

type Followers struct {
	gorm.Model `json:"-"`
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	UserName string `json:"userName,omitempty"`

	Following []*Following `gorm:"many2many:following_followers;" json:"-"`

	UserId  int64  `json:"userID,omitempty"`
}
