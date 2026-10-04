package models

import "time"

type LoginVerification struct {
	ID        string    `gorm:"column:loginverificationid;type:uuid;primaryKey;default:gen_random_uuid()" json:"-"`
	UserID    int       `gorm:"column:userid;not null;index" json:"-"`
	Token     string    `gorm:"column:token;type:text;not null;unique" json:"-"`
	ExpiresAt time.Time `gorm:"column:expiresat;not null" json:"-"`
	Verified  bool      `gorm:"column:verified;not null;default:false" json:"-"`
	CreatedAt time.Time `gorm:"column:createdat;not null;default:CURRENT_TIMESTAMP" json:"-"`
}

func (LoginVerification) TableName() string {
	return "loginverifications"
}
