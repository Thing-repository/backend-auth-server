package moduleErrors

import "errors"

var (
	ErrorServiceUserNotFound    = errors.New("user not found")
	ErrorServiceInvalidPassword = errors.New("invalid password")
	ErrorServiceUserAlreadyHas  = errors.New("there is already a user with this email")
)
