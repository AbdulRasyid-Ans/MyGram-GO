package params

import "time"

type UserSocialMedia struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type SocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"UserId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	User           UserSocialMedia
}

type RequestCreateOrUpdateSocialMedia struct {
	Name           string `json:"name" valid:"required~Name tidak boleh kosong"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~Social Media URL tidak boleh kosong"`
}

type ResponseCreateSocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type ResponseUpdateSocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ResponseGetSocialMedia struct {
	SocialMedias []SocialMedia `json:"social_medias"`
}
