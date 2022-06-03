package core

type UserSignInData struct {
	UserMail     string `json:"email"`
	UserPassword string `json:"password"`
}

type UserSignUpData struct {
	UserBaseData
	Password string `json:"password" binding:"required"`
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

type SignInResponse struct {
	User
	Token string `json:"token"`
}
