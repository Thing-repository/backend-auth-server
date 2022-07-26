package core

type CompanyBase struct {
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type Company struct {
	CompanyBase
	ImageURL *string `json:"image_url,omitempty"`
	Id       int     `json:"id"`
}
