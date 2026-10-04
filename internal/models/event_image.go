package models

type EventImage struct {
	EventImageID uint   `gorm:"column:eventimagesid;primaryKey"`
	EventID      uint   `gorm:"column:eventid;not null"`
	ImageURL     string `gorm:"column:imageurl;size:500;not null"`
	DisplayOrder int    `gorm:"column:displayorder;not null"`

	Event Event `gorm:"foreignKey:EventID"`
}

func (EventImage) TableName() string {
	return "eventimages"
}
