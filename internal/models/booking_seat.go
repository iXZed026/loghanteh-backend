package models

type BookingSeat struct {
	BookingSeatID uint `gorm:"column:bookingseatid;primaryKey"`

	BookingID uint `gorm:"column:bookingid;not null"`
	SessionID uint `gorm:"column:sessionid;not null"`
	SeatID    uint `gorm:"column:seatid;not null"`

	Booking Booking      `gorm:"foreignKey:BookingID"`
	Session EventSession `gorm:"foreignKey:SessionID"`
	Seat    Seat         `gorm:"foreignKey:SeatID"`
}

func (BookingSeat) TableName() string {
	return "bookingseats"
}
