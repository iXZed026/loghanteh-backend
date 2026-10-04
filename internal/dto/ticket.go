package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

// --------------------------------------------------
// Museum Tour
// --------------------------------------------------

type MuseumTourSessionResponse struct {
	SessionID   uint      `json:"sessionId"`
	EventID     uint      `json:"eventId"`
	HallID      uint      `json:"hallId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"startAt"`
	Duration    int       `json:"duration"`
	Capacity    int       `json:"capacity"`
	Price       *float64  `json:"price"`
}

type MuseumTourQuery struct {
	Date string `form:"date" binding:"required"`
}

// --------------------------------------------------
// Events / Courses
// --------------------------------------------------

type EventSessionListResponse struct {
	SessionID uint `json:"sessionId"`
	EventID   uint `json:"eventId"`

	Name   string               `json:"name"`
	Images []EventImageResponse `json:"images"`

	StartAt time.Time `json:"startAt"`
	Price   *float64  `json:"price"`
}

type EventImageResponse struct {
	ImageURL     string `json:"imageUrl"`
	DisplayOrder int    `json:"displayOrder"`
}

type EventQuery struct {
	Date string `form:"date" binding:"required"`
}

type CourseQuery struct {
	Date string `form:"date" binding:"required"`
}

// --------------------------------------------------
// Cinema
// --------------------------------------------------

type CinemaQuery struct {
	Date string `form:"date" binding:"required"`
}

type CinemaSessionListResponse struct {
	SessionID     uint                  `json:"sessionId"`
	EventID       uint                  `json:"eventId"`
	HallID        uint                  `json:"hallId"`
	Name          string                `json:"name"`
	Description   string                `json:"description"`
	Images        []EventImageResponse  `json:"images"`
	StartAt       time.Time             `json:"startAt"`
	Duration      int                   `json:"duration"`
	Price         *float64              `json:"price"`
	CinemaDetails *CinemaDetailResponse `json:"cinemaDetails"`
}

type CinemaDetailResponse struct {
	EventID      uint     `json:"eventId"`
	ReleaseYear  *int     `json:"releaseYear"`
	Director     *string  `json:"director"`
	Country      *string  `json:"country"`
	FilmDuration *int     `json:"filmDuration"`
	Genre        *string  `json:"genre"`
	IMDBScore    *float64 `json:"imdbScore"`
}

// --------------------------------------------------
// Theater
// --------------------------------------------------

type TheaterQuery struct {
	Date string `form:"date" binding:"required"`
}

type TheaterSessionListResponse struct {
	SessionID      uint                   `json:"sessionId"`
	EventID        uint                   `json:"eventId"`
	HallID         uint                   `json:"hallId"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	Images         []EventImageResponse   `json:"images"`
	StartAt        time.Time              `json:"startAt"`
	Duration       int                    `json:"duration"`
	Price          *float64               `json:"price"`
	TheaterDetails *TheaterDetailResponse `json:"theaterDetails"`
}

type TheaterDetailResponse struct {
	EventID  uint   `json:"eventId"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Duration int    `json:"duration"`
}

// --------------------------------------------------
// Event detail
// --------------------------------------------------

type EventSessionDetailResponse struct {
	SessionID   uint                 `json:"sessionId"`
	EventID     uint                 `json:"eventId"`
	HallID      uint                 `json:"hallId"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Images      []EventImageResponse `json:"images"`
	StartAt     time.Time            `json:"startAt"`
	Duration    int                  `json:"duration"`
	Price       *float64             `json:"price"`
}

// --------------------------------------------------
// Booking Seats
// --------------------------------------------------

type SessionSeatResponse struct {
	SeatID     uint   `json:"seatId"`
	RowLabel   string `json:"rowLabel"`
	SeatNumber int    `json:"seatNumber"`
	Status     string `json:"status"`
}

type CreateSeatBookingRequest struct {
	SessionID uint   `json:"sessionId" binding:"required"`
	SeatIDs   []uint `json:"seatIds" binding:"required,min=1"`
}

type UserReservedSessionResponse struct {
	BookingID   uint `json:"bookingId"`
	SessionID   uint `json:"sessionId"`
	EventID     uint `json:"eventId"`
	EventTypeID uint `json:"eventTypeId"`
	HallID      uint `json:"hallId"`

	Quantity   int             `json:"quantity"`
	TotalPrice decimal.Decimal `json:"totalPrice"`

	HallName      string `json:"hallName"`
	EventTypeName string `json:"eventTypeName"`

	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"startAt"`
	Duration    int       `json:"duration"`
	Price       *float64  `json:"price"`

	CinemaDetails  *CinemaDetailResponse  `json:"cinemaDetails,omitempty"`
	TheaterDetails *TheaterDetailResponse `json:"theaterDetails,omitempty"`

	Seats []SessionSeatResponse `json:"seats"`
}
