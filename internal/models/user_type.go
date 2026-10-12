package models

type UserType struct {
	UserTypeID  int    `json:"userTypeId" gorm:"column:usertypeid;primaryKey"`
	Code        string `json:"code" gorm:"column:code;size:30"`
	Name        string `json:"name" gorm:"column:name;size:100"`
	Description string `json:"description" gorm:"column:description;type:text"`
}

func (UserType) TableName() string {
	return "usertypes"
}
