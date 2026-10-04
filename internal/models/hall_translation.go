package models

type HallTranslation struct {
	HallID      uint   `gorm:"column:hallid;primaryKey"`
	LanguagesID uint   `gorm:"column:languagesid;primaryKey"`
	Name        string `gorm:"column:name;size:255;not null"`
	Description string `gorm:"column:description;type:text"`

	Hall Hall `gorm:"foreignKey:HallID"`
}

func (HallTranslation) TableName() string {
	return "halltranslations"
}
