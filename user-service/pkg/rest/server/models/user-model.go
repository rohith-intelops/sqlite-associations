package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Id         int64 `gorm:"primaryKey;autoIncrement" json:"ID,omitempty"`

	EmailId string `json:"emailId,omitempty"`

	Name string `json:"name,omitempty"`

	Post       Post       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserId"`

	Followers  Followers `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserId"` 

	Following  Following `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:UserId"`
}
