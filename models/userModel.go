package models

import (
	"final-project/helpers"

	"gorm.io/gorm"
)

type User struct {
	Model
	Username     string        `gorm:"type:varchar(255);not null"`
	Email        string        `gorm:"type:varchar(255);not null;uniqueIndex"`
	Password     string        `gorm:"not null"`
	Age          int           `gorm:"type:integer;not null"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
