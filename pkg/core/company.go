package core

type CompanyBase struct {
	CompanyName string `json:"company_name" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type Company struct {
	CompanyBase
	ImageURL string `json:"image_url"`
	Id       int    `json:"id"`
}
