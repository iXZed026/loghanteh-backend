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
	PhoneNumber     string `json:"phoneNumber" binding:"omitempty,min=11,max=20"`
	Email           string `json:"email" binding:"omitempty,email,min=3,max=254"`
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
