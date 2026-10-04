package models

import "time"

type RegisterVerification struct {
	RegisterVerificationID string `gorm:"type:uuid;primaryKey;column:registerverificationid;default:gen_random_uuid()" json:"-"`

	FullName     string `gorm:"column:fullname;type:varchar(150);not null" json:"-"`
	PhoneNumber  string `gorm:"column:phonenumber;type:varchar(20);not null;unique" json:"-"`
	Email        string `gorm:"column:email;type:varchar(254);not null;unique" json:"-"`
	PasswordHash string `gorm:"column:passwordhash;type:text;not null" json:"-"`

	Token     string    `gorm:"column:token;type:text;not null;unique" json:"-"`
	ExpiresAt time.Time `gorm:"column:expiresat;type:timestamp with time zone;not null" json:"-"`
	Verified  bool      `gorm:"column:verified;type:boolean;not null;default:false" json:"-"`
	CreatedAt time.Time `gorm:"column:createdat;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP" json:"-"`
}

func (RegisterVerification) TableName() string {
	return "registerverifications"
}
