package models

type CinemaDetail struct {
	EventID      uint     `gorm:"column:eventid;primaryKey"`
	ReleaseYear  *int     `gorm:"column:releaseyear"`
	Director     *string  `gorm:"column:director"`
	Country      *string  `gorm:"column:country"`
	FilmDuration *int     `gorm:"column:filmduration"`
	Genre        *string  `gorm:"column:genre"`
	IMDBScore    *float64 `gorm:"column:imdbscore"`
}

func (CinemaDetail) TableName() string {
	return "cinemadetails"
}
