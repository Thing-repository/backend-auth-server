package core

type User struct {
	Id                int     `json:"id" binding:"required"`
	FirstName         string  `json:"first_name" binding:"required"`
	LastName          string  `json:"last_name" binding:"required"`
	Email             string  `json:"email" binding:"required"`
	ImageURL          string  `json:"image_url" binding:"required"`
	CompanyId         *int    `json:"company_id,omitempty"`
	DepartmentId      *int    `json:"department_id,omitempty"`
	IsCompanyAdmin    *bool   `json:"is_company_admin,omitempty"`
	IsDepartmentAdmin *bool   `json:"is_department_admin,omitempty"`
	VacationTimeStart *uint32 `json:"vacation_time_start,omitempty"`
	VacationTimeEnd   *uint32 `json:"vacation_time_end,omitempty"`
}
