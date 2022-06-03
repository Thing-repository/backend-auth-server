package handler

import (
	_ "github.com/Thing-repository/backend-server/docs"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	api := router.Group("/api/v1")
	{
		api.GET("/open_api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", H.signUp)
			auth.POST("/sign-in", H.signIn)
		}
	}
	return router
}
