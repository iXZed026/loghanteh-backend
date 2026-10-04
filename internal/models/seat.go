package models

type Seat struct {
	SeatID     uint   `gorm:"column:seatid;primaryKey"`
	HallID     uint   `gorm:"column:hallid;not null"`
	RowLabel   string `gorm:"column:rowlabel;size:10;not null"`
	SeatNumber int    `gorm:"column:seatnumber;not null"`

	Hall Hall `gorm:"foreignKey:HallID"`
}

func (Seat) TableName() string {
	return "seats"
}
