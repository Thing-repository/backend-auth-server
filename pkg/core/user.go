package core

type UserBaseData struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type UserChange struct {
	UserBaseData
	VacationTimeStart *uint32 `json:"vacation_time_start,omitempty"`
	VacationTimeEnd   *uint32 `json:"vacation_time_end,omitempty"`
}

type User struct {
	UserChange
	EmailIsValidated bool    `json:"email_is_validated"`
	Id               int     `json:"id" binding:"required"`
	ImageURL         *string `json:"image_url,omitempty"`
	CompanyId        *int    `json:"company_id,omitempty"`
	DepartmentId     *int    `json:"department_id,omitempty"`
}
