package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Booking struct {
	BookingID uint `gorm:"column:bookingid;primaryKey"`

	UserID    uint `gorm:"column:userid;not null"`
	SessionID uint `gorm:"column:sessionid;not null"`

	Quantity   int             `gorm:"column:quantity;not null"`
	TotalPrice decimal.Decimal `gorm:"column:totalprice;type:numeric(12,2);not null"`

	PurchasedAt time.Time `gorm:"column:purchasedat;not null;default:now()"`

	User    User          `gorm:"foreignKey:UserID"`
	Session EventSession  `gorm:"foreignKey:SessionID"`
	Seats   []BookingSeat `gorm:"foreignKey:BookingID"`
}

func (Booking) TableName() string {
	return "bookings"
}
