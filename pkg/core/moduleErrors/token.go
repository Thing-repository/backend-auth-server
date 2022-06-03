package moduleErrors

import "errors"

var (
	ErrorTokenInvalidToken = errors.New("error invalid token")
	ErrorTokenExpiredToken = errors.New("error expired token")
)
