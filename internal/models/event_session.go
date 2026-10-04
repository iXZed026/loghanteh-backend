package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type EventSession struct {
	SessionID uint `gorm:"column:sessionid;primaryKey"`
	EventID   uint `gorm:"column:eventid;not null"`
	HallID    uint `gorm:"column:hallid;not null"`

	StartAt  time.Time        `gorm:"column:startat;not null"`
	Duration int              `gorm:"column:duration;not null"`
	Capacity int              `gorm:"column:capacity;not null"`
	Price    *decimal.Decimal `gorm:"column:price;type:numeric(12,2)"`

	Event Event `gorm:"foreignKey:EventID"`
	Hall  Hall  `gorm:"foreignKey:HallID"`
}

func (EventSession) TableName() string {
	return "eventsession"
}
