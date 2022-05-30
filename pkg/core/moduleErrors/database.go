package moduleErrors

import "errors"

var (
	ErrorDatabaseUserNotFound   = errors.New("user not found")
	ErrorDataBaseInternal       = errors.New("internal data base error")
	ErrorDataBaseUserAlreadyHas = errors.New("there is already a user with this email")
)
