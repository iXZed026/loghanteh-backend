package dto

// Register
type RegisterRequest struct {
	FullName       string `json:"fullName" binding:"required,max=150"`
	PhoneNumber    string `json:"phoneNumber" binding:"omitempty,min=11,max=20"`
	Email          string `json:"email" binding:"omitempty,email,min=3,max=254"`
	Password       string `json:"password" binding:"required,min=8"`
	RepeatPassword string `json:"repeatPassword" binding:"required"`
}

type RegisterResponse struct {
	UserID            int    `json:"userId"`
	FullName          string `json:"fullName"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	VerificationToken string `json:"verificationToken"`
}

// Login
type LoginRequest struct {
	EmailOrPhone string `json:"emailOrPhone" binding:"required,max=254"`
	Password     string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	UserID            int    `json:"userId"`
	FullName          string `json:"fullName"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	VerificationToken string `json:"verificationToken"`
}

// Forget Password
type ForgetPasswordRequest struct {
	EmailOrPhone string `json:"emailOrPhone" binding:"required,max=254"`
}

type ForgetPasswordResponse struct {
	UserID            int    `json:"userId"`
	FullName          string `json:"fullName"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	VerificationToken string `json:"verificationToken"`
}
