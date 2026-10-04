package dto

type NewPasswordRequest struct {
	VerificationToken string `json:"verificationToken" binding:"required"`
	Password          string `json:"password" binding:"required,min=8"`
	RepeatPassword    string `json:"repeatPassword" binding:"required,min=8"`
}

type NewPasswordResponse struct {
	UserID uint `json:"userId"`
}
