package moduleErrors

import "errors"

var (
	ErrorHandlerInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrorHandlerInvalidEmail              = errors.New("invalid email address")
	ErrorHandlerNoRequiredFieldsQuery     = errors.New("no required Fields in query")
	ErrorHandlerForbidden                 = errors.New("you have not access to this resource")
)
