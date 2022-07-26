package handler

import (
	"context"
	"errors"
	_ "github.com/Thing-repository/backend-server/docs"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//go:generate mockgen -source=handler.go -destination=mocks/authMock.go
type auth interface {
	SignIn(authData *core.UserSignInData) (*core.SignInResponse, error)
	SignUp(authData *core.UserSignUpData) (*core.SignInResponse, error)
}

//go:generate mockgen -source=handler.go -destination=mock/authMock.go
type company interface {
	AddCompany(companyAdd *core.CompanyBase, user *core.User) (*core.Company, error)
	GetCompany(companyId int) (*core.Company, error)
	UpdateCompany(companyBase core.CompanyBase, companyId int) (*core.Company, error)
	DeleteCompany(companyId int) error
}

//go:generate mockgen -source=handler.go -destination=mock/authMock.go
type token interface {
	ValidateToken(token string) (int, error)
}

//go:generate mockgen -source=auth.go -destination=mock/authMock.go
type userDB interface {
	GetUser(ctx context.Context, userId int) (*core.UserDB, error)
	//UserIsCompanyAdmin(userId int, companyId int) (bool, error)
	//UserIsDepartmentAdmin(userId int, departmentId int) (bool, error)
	//UserIsDepartmentMaintainer(userId int, departmentId int) (bool, error)
}

type Handler struct {
	auth    auth
	token   token
	company company
	userDB  userDB
}

func NewHandler(auth auth, company company, token token, userDB userDB) *Handler {
	return &Handler{
		auth:    auth,
		token:   token,
		company: company,
		userDB:  userDB,
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
	apiPrivate := router.Group("/api/v1", H.userIdentity)
	{
		company := apiPrivate.Group("/company")
		{
			company.POST("", H.addCompany)
			company.GET("/:company_id", H.getCompany)
			company.PATCH("/:company_id", H.patchCompany)
			company.DELETE("/:company_id", H.deleteCompany)
		}
	}
	return router
}

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("can't get user id")
	}
	id, ok := userId.(int)
	if !ok {
		return 0, errors.New("can't get user id")
	}
	return id, nil
}
