package core

type DepartmentBase struct {
	DepartmentName string `json:"department_name" binding:"required"`
}

type Department struct {
	DepartmentBase
	Id       int    `json:"id"`
	ImageURL string `json:"image_url"`
}
