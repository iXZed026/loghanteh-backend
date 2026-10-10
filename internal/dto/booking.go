package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateBookingRequest struct {
	SessionID      uint  `json:"sessionId" binding:"required"`
	Quantity       int   `json:"quantity" binding:"required,min=1"`
	DiscountCodeID *uint `json:"discountCodeId"`
}

type CreateBookingResponse struct {
	BookingID   uint            `json:"bookingId"`
	SessionID   uint            `json:"sessionId"`
	Quantity    int             `json:"quantity"`
	TotalPrice  decimal.Decimal `json:"totalPrice"`
	PurchasedAt time.Time       `json:"purchasedAt"`
}

type BookingCountResponse struct {
	Count int64 `json:"count"`
}

type NextBookingResponse struct {
	BookingID   uint      `json:"bookingId"`
	Name        string    `json:"name"`
	TicketToken uuid.UUID `json:"ticketToken"`
	StartAt     time.Time `json:"startAt"`
}

type UserBookingResponse struct {
	BookingID     uint   `json:"bookingId"`
	EventTypeID   uint   `json:"eventTypeId"`
	EventTypeName string `json:"eventTypeName"`
	EventID       uint   `json:"eventId"`
	SessionID     uint   `json:"sessionId"`

	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"startAt"`
	Duration    int       `json:"duration"`

	Quantity    int             `json:"quantity"`
	TotalPrice  decimal.Decimal `json:"totalPrice"`
	PurchasedAt time.Time       `json:"purchasedAt"`

	DiscountCodeID *uint           `json:"discountCodeId"`
	DiscountAmount decimal.Decimal `json:"discountAmount"`
	TicketToken    uuid.UUID       `json:"ticketToken"`
	TicketIsValid  bool            `json:"ticketIsValid"`
}
