package handler

import (
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=mocks/authMock.go
type Auth interface {
	SignIn(authData *core.UserSignInData) (*core.SignInResponse, error)
	SignUp(authData *core.UserSignUpData) (*core.SignInResponse, error)
}

type Handler struct {
	auth Auth
}

func NewHandler(auth Auth) *Handler {
	return &Handler{
		auth: auth,
	}
}

func (H *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Group("/api/v1")
	{
		auth := router.Group("/auth")
		{
			auth.POST("/sign-up", H.signUp)
			auth.POST("/sign-in", H.signIn)
		}
	}
	return router
}
