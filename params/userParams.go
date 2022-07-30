package params

import (
	"time"
)

type RequestRegisterUser struct {
	Age      int    `json:"age" valid:"required~Age tidak boleh kosong"`
	Email    string `json:"email" valid:"required~Email tidak boleh kosong, email~Format email salah"`
	Password string `json:"password" valid:"required~Password tidak boleh kosong, minstringlength(6)~Password minimal 6 karakter"`
	Username string `json:"username" valid:"required~Username tidak boleh kosong"`
}

type RequestUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestUpdateUser struct {
	Email    string `json:"email" valid:"required~Email tidak boleh kosong, email~Format email salah"`
	Username string `json:"username" valid:"required~Username tidak boleh kosong"`
}

type ResponseRegisterUser struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type ResponseUpdateUser struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}
