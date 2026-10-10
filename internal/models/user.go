package models

import "time"

type User struct {
	UserID       int        `json:"user_Id" gorm:"column:userid;primaryKey"`
	FullName     string     `json:"full_name" gorm:"column:fullname"`
	Email        *string    `json:"email" gorm:"column:email"`
	PhoneNumber  *string    `json:"phone_number" gorm:"column:phonenumber"`
	DOB          *time.Time `json:"dob,omitempty" gorm:"column:dob"`
	PasswordHash string     `json:"-" gorm:"column:passwordhash"`
}

func (User) TableName() string {
	return "users"
}
