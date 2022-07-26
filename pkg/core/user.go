package core

type UserBaseData struct {
	FirstName *string `json:"first_name" binding:"required"`
	LastName  *string `json:"last_name" binding:"required"`
	Email     *string `json:"email" binding:"required"`
}

type User struct {
	UserBaseData
	EmailIsValidated *bool   `json:"email_is_validated"`
	Id               int     `json:"id" binding:"required"`
	ImageURL         *string `json:"image_url,omitempty"`
	CompanyId        *int    `json:"company_id,omitempty"`
	DepartmentId     *int    `json:"department_id,omitempty"`
}

type UserDB struct {
	User
	PasswordHash         *string `json:"password_hash"`
	EmailValidationToken *string `json:"email_validation_token"`
}

type AddUserDB struct {
	UserBaseData
	PasswordHash string `json:"password_hash"`
}
