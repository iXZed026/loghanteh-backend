package models

type Event struct {
	EventID     uint `gorm:"column:eventid;primaryKey"`
	EventTypeID uint `gorm:"column:eventtypeid;not null"`

	EventType EventType      `gorm:"foreignKey:EventTypeID"`
	Sessions  []EventSession `gorm:"foreignKey:EventID"`
	Images    []EventImage   `gorm:"foreignKey:EventID"`
}

func (Event) TableName() string {
	return "events"
}
