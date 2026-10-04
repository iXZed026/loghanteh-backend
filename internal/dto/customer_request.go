package dto

type CustomerRequestReq struct {
	FullName     string `json:"fullName" binding:"required,max=150"`
	EmailOrPhone string `json:"emailOrPhone" binding:"required,max=255"`
	CompanyName  string `json:"companyName" binding:"max=255"`
	RequestTitle string `json:"requestTitle" binding:"required,max=255"`
	Description  string `json:"description" binding:"max=500"`
}
