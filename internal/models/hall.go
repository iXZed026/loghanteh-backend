package models

type Hall struct {
	HallID uint `gorm:"column:hallid;primaryKey"`

	Sessions []EventSession `gorm:"foreignKey:HallID"`
	Seats    []Seat         `gorm:"foreignKey:HallID"`
}

func (Hall) TableName() string {
	return "halls"
}
