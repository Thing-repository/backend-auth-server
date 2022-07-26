package moduleErrors

import "errors"

var (
	ErrorHandlerInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrorHandlerInvalidEmail              = errors.New("invalid email address")
	ErrorHandlerUserAlreadyHasCompany     = errors.New("user already has company")
	ErrorNoRequiredFieldsQuery            = errors.New("no required Fields in query")
	ErrorForbidden                        = errors.New("you have not access to this resource")
)
