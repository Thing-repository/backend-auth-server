package core

type UserDB struct {
	User
	PasswordHash string `json:"password_hash"`
}

type AddUserDB struct {
	UserBaseData
	PasswordHash string `json:"password_hash"`
}
