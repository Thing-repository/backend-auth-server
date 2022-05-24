package rest_handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
