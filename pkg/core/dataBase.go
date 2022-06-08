package core

type UserDB struct {
	User
	PasswordHash         string `json:"password_hash"`
	EmailValidationToken string `json:"email_validation_token"`
}

type AddUserDB struct {
	UserBaseData
	PasswordHash string `json:"password_hash"`
}
