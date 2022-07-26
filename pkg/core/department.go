package core

type DepartmentBase struct {
	DepartmentName string `json:"department_name" binding:"required"`
	CompanyId      int    `json:"company_id" binding:"required"`
}

type Department struct {
	DepartmentBase
	Id       int    `json:"id"`
	ImageURL string `json:"image_url"`
}
