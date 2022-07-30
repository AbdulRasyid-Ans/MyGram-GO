package params

import "time"

type UserPhoto struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type RequestCreateOrUpdatePhoto struct {
	Title    string `json:"title" valid:"required~Title tidak boleh kosong"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo URL tidak boleh kosong"`
}

type ResponseCreatePhoto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseUpdatePhoto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseGetPhoto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserPhoto
}
