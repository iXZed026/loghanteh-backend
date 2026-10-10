package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type DiscountCode struct {
	DiscountCodeID    uint            `gorm:"column:discountcodeid;primaryKey"`
	Code              string          `gorm:"column:code;size:50;not null"`
	DiscountPercent   decimal.Decimal `gorm:"column:discountpercent;type:numeric(5,2);not null"`
	MaxDiscountAmount decimal.Decimal `gorm:"column:maxdiscountamount;type:numeric(12,2);not null"`
	UsageLimit        int             `gorm:"column:usagelimit;not null"`
	UsedCount         int             `gorm:"column:usedcount;not null"`
	StartAt           time.Time       `gorm:"column:startat;not null"`
	ExpiresAt         time.Time       `gorm:"column:expiresat;not null"`
	IsActive          bool            `gorm:"column:isactive;not null"`
	CreatedAt         time.Time       `gorm:"column:createdat;not null"`
}

func (DiscountCode) TableName() string {
	return "discountcodes"
}
