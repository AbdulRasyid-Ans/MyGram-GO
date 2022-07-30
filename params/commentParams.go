package params

import "time"

type PhotoComment struct {
	ID       uint   `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Caption  string `json:"caption,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	UserID   uint   `json:"user_id,omitempty"`
}

type UserComment struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type RequestCreateComment struct {
	Message string `json:"message" valid:"required~Message tidak boleh kosong"`
	PhotoID uint   `json:"photo_id"`
}

type RequestUpdateComment struct {
	Message string `json:"message" valid:"required~Message tidak boleh kosong"`
}

type ResponseCreateComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseUpdateComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseGetComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      UserComment
	Photo     PhotoComment
}
