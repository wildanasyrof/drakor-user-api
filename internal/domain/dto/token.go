package dto

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
	AccessExp   int64  `json:"access_exp"`
}
