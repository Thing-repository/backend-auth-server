package core

import (
	"golang.org/x/exp/slices"
)

const (
	CredentialTypeCompanyAdmin         = "company_admin"
	CredentialTypeDepartmentAdmin      = "department_admin"
	CredentialTypeDepartmentMaintainer = "department_maintainer"
	CredentialTypeCompanyUser          = "company_user"
	CredentialTypeDepartmentUser       = "department_user"
)

var CompanyCredential = []string{CredentialTypeCompanyAdmin, CredentialTypeCompanyUser}
var DepartmentCredential = []string{CredentialTypeDepartmentAdmin, CredentialTypeDepartmentMaintainer, CredentialTypeDepartmentUser}

type Credentials struct {
	CredentialType string
	ObjectId       int
}

type AddCredentials struct {
	Credentials
	UserId int
}

type CredentialsDb struct {
	Id int
	AddCredentials
	Credentials
}

func CheckCredential(credentials []Credentials, CredentialType string, ObjectId int) bool {
	idx := slices.IndexFunc(credentials, func(c Credentials) bool {
		return c.CredentialType == CredentialType && c.ObjectId == ObjectId
	})
	return idx != -1
}
