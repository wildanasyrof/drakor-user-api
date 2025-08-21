package dto

import "github.com/google/uuid"

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	ImgUrl   string    `json:"img_url"`
}
