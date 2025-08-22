package dto

type UpdateUserRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	Username *string `json:"username" validate:"omitempty,min=3,max=32"`
	Password *string `json:"password" validate:"omitempty,min=6"`
	ImgUrl   *string `json:"img_url" validate:"omitempty,url"`
}
