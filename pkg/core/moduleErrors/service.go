package moduleErrors

import "errors"

var (
	ErrorServiceUserNotFound          = errors.New("user not found")
	ErrorServiceInvalidPassword       = errors.New("invalid password")
	ErrorServiceUserAlreadyHas        = errors.New("there is already a user with this email")
	ErrorServiceGetUserData           = errors.New("error get user data")
	ErrorServiceUserAlreadyHasCompany = errors.New("user already has company")
	ErrorServiceBadPermissions        = errors.New("access denied")
	ErrorServiceInvalidContext        = errors.New("invalid context")
	ErrorAllNoFields                  = errors.New("nothing to change")
)
