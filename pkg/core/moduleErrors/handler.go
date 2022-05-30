package moduleErrors

import "errors"

var (
	ErrorHandlerInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrorInvalidEmail                     = errors.New("invalid email address")
)
