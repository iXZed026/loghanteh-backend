package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

// --------------------------------------------------
// Discount Code
// --------------------------------------------------

type DiscountCodeRequest struct {
	Code string `json:"code" binding:"required,max=50"`
}

type DiscountCodeResponse struct {
	DiscountCodeID    uint            `json:"discountCodeId"`
	Code              string          `json:"code"`
	DiscountPercent   decimal.Decimal `json:"discountPercent"`
	MaxDiscountAmount decimal.Decimal `json:"maxDiscountAmount"`
	UsageLimit        int             `json:"usageLimit"`
	UsedCount         int             `json:"usedCount"`
	StartAt           time.Time       `json:"startAt"`
	ExpiresAt         time.Time       `json:"expiresAt"`
	IsActive          bool            `json:"isActive"`
	CreatedAt         time.Time       `json:"createdAt"`
}
