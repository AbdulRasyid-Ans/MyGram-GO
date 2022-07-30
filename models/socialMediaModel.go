package models

type SocialMedia struct {
	Model
	Name           string `gorm:"type:varchar(255);not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserID         uint

	User *User
}
