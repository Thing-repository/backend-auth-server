package moduleErrors

import "errors"

var (
	ErrorHandlerInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrorHandlerInvalidEmail              = errors.New("invalid email address")
	ErrorHandlerUserAlreadyHasCompany     = errors.New("user already has company")
)
