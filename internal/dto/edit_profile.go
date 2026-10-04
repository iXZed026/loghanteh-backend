package dto

import "time"

type ProfileResponse struct {
	UserID      uint       `json:"userId"`
	FullName    string     `json:"fullName"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phoneNumber"`
	DOB         *time.Time `json:"dob"`
}

type EditProfileRequest struct {
	FullName        string `json:"fullName" binding:"required,max=150"`
	Email           string `json:"email" binding:"required,email,max=254"`
	PhoneNumber     string `json:"phoneNumber" binding:"required,max=20"`
	DOB             string `json:"dob"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword" binding:"omitempty,min=8"`
}

type EditProfileResponse struct {
	UserID      uint       `json:"userId"`
	FullName    string     `json:"fullName"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phoneNumber"`
	DOB         *time.Time `json:"dob"`
}
