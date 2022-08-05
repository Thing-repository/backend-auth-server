package moduleErrors

import "errors"

var (
	ErrorDatabaseUserNotFound           = errors.New("user not found")
	ErrorDataBaseInternal               = errors.New("internal data base error")
	ErrorDataBaseUserAlreadyHas         = errors.New("there is already a user with this email")
	ErrorDataBaseGetTransaction         = errors.New("error get transaction")
	ErrorDataBaseHasNotDataToChange     = errors.New("error hasn't data to change")
	ErrorDataBaseInvalidCredentialsType = errors.New("invalid credentials type")
)
