package models

type Comment struct {
	Model
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null"`

	User  *User
	Photo *Photo
}
