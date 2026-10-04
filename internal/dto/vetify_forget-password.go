package dto

type VerifyForgetPasswordRequest struct {
	VerificationToken string `json:"verificationToken" binding:"required"`
	Code              string `json:"code" binding:"required"`
}

type VerifyForgetPasswordResponse struct {
	UserID      int    `json:"userId"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
