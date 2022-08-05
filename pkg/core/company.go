package core

type CompanyBase struct {
	CompanyName *string `json:"company_name,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type CompanyUpdate struct {
	CompanyBase
	ImageURL *string `json:"image_url,omitempty"`
}

type Company struct {
	CompanyUpdate
	Id int `json:"id"`
}
