package models

type EventType struct {
	EventTypeID   uint   `gorm:"column:eventtypeid;primaryKey"`
	EventTypeName string `gorm:"column:eventtypename;size:100;not null"`

	Events []Event `gorm:"foreignKey:EventTypeID"`
}

func (EventType) TableName() string {
	return "eventtypes"
}
