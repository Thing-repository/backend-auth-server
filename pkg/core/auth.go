package core

type UserAuthData struct {
	UserMail     string `json:"user_mail"`
	UserPassword string `json:"user_password"`
}

type Access struct {
	CompanyId         int  `json:"company_id"`
	IsCompanyAdmin    bool `json:"is_company_admin"`
	DepartmentId      int  `json:"department_id"`
	IsDepartmentAdmin bool `json:"is_department_admin"`
}

type AuthValidationData struct {
	UserId       int    `json:"user_id"`
	UserAccesses Access `json:"user_accesses"`
}
