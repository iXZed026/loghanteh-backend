package models

type UserRole struct {
	UserRoleID  int    `json:"userRoleId" gorm:"column:userroleid;primaryKey"`
	Code        string `json:"code" gorm:"column:code;size:30"`
	Name        string `json:"name" gorm:"column:name;size:100"`
	Description string `json:"description" gorm:"column:description;type:text"`
}

func (UserRole) TableName() string {
	return "userroles"
}
