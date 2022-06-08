package core

type CompanyManager struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	CompanyId int `json:"company_id"`
}

type DepartmentManager struct {
	Id           int `json:"id"`
	UserId       int `json:"user_id"`
	DepartmentId int `json:"department_id"`
}
