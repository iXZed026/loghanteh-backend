package models

type CustomerRequest struct {
	RequestID    int    `gorm:"column:requestid;primaryKey"`
	FullName     string `gorm:"column:fullname"`
	CompanyName  string `gorm:"column:companyname"`
	Email        string `gorm:"column:email"`
	PhoneNumber  string `gorm:"column:phonenumber"`
	RequestTitle string `gorm:"column:requesttitle"`
	Description  string `gorm:"column:description"`
}

func (CustomerRequest) TableName() string {
	return "customerrequests"
}
