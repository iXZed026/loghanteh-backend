package models

import "time"

// type RefreshToken struct {
// 	RefreshTokenUUID string    `gorm:"type:uuid;primaryKey;column:refreshtokenid;default:gen_random_uuid()" json:"-"`
// 	UserID           int       `gorm:"column:userid;not null;index" json:"-"`
// 	Token            string    `gorm:"column:token;type:text;not null;unique" json:"-"`
// 	ExpiresAt        time.Time `gorm:"column:expiresat;type:timestamp with time zone;not null" json:"-"`
// 	Revoked          bool      `gorm:"column:revoked;type:boolean;not null;default:false" json:"-"`
// 	CreatedAt        time.Time `gorm:"column:createdat;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP" json:"-"`
// }
//test
type RefreshToken struct {
	RefreshTokenUUID string    `gorm:"type:uuid;primaryKey;column:refreshtokenid;default:gen_random_uuid()" json:"refreshTokenId"`
	UserID           int       `gorm:"column:userid;not null;index" json:"ser_id"`
	Token            string    `gorm:"column:token;type:text;not null;unique" json:"token"`
	ExpiresAt        time.Time `gorm:"column:expiresat;type:timestamp with time zone;not null" json:"exoireAt"`
	Revoked          bool      `gorm:"column:revoked;type:boolean;not null;default:false" json:"revoked"`
	CreatedAt        time.Time `gorm:"column:createdat;type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
}

func (RefreshToken) TableName() string {
	return "refreshtoken"
}
