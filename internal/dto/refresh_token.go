package dto

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type RefreshResponse struct {
	AccessToken string `json:"accessToken"`
}
