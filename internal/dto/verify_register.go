package dto

type VerifyRegisterRequest struct {
	VerificationToken string `json:"verificationToken" binding:"required"`
	Code              string `json:"code" binding:"required"`
}

type VerifyRegisterResponse struct {
	UserID      int    `json:"userId"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	AccessToken string `json:"accessToken"`
}
