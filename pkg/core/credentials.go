package core

const (
	CredentialTypeCompanyAdmin         = "company_admin"
	CredentialTypeDepartmentAdmin      = "department_admin"
	CredentialTypeDepartmentMaintainer = "department_maintainer"
)

type Credentials struct {
	CredentialType string
	UserId         int
	ObjectId       int
}

type CredentialsDb struct {
	Id int
	Credentials
}
