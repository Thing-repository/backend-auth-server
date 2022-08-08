package handler

import (
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary NOT IMPLEMENTED! CurrentUser
// @Security ApiKeyAuth
// @Tags user
// @Description This request for get current user info
// @ID getCurrentUser
// @Accept json
// @Produces json
// @Success 200 {object} core.User
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [get]
func (H *Handler) getCurrentUser(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! User
// @Security ApiKeyAuth
// @Tags user
// @Description This request for get user info
// @ID getUser
// @Accept json
// @Produces json
// @Param id path int true "user id"
// @Success 200 {object} core.User
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/find [get]
func (H *Handler) getUser(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! FindUsersForInvite
// @Security ApiKeyAuth
// @Tags user
// @Description This request for get user info
// @ID findUsersForInvite
// @Accept json
// @Produces json
// @Param filter query string true "filter for find"
// @Param limit query int true "limit for found users list"
// @Param offset query int true "offset for found users list"
// @Success 200 {array} core.User
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/find [get]
func (H *Handler) findUsersForInvite(c *gin.Context) {
	var filter string
	var limit string
	var offset string

	var limitInt int
	var offsetInt int

	var ok bool
	var err error

	if filter, ok = c.GetQuery("filter"); !ok {
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		return
	}
	if limit, ok = c.GetQuery("limit"); !ok {
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		return
	}
	if offset, ok = c.GetQuery("offset"); !ok {
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		return
	}
	if limitInt, err = strconv.Atoi(limit); err != nil {
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		return
	}
	if offsetInt, err = strconv.Atoi(offset); err != nil {
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		return
	}
	users, err := H.user.FindUsersForInvite(c, filter, limitInt, offsetInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.AbortWithStatusJSON(http.StatusOK, users)
}

// @Summary NOT IMPLEMENTED! Users
// @Security ApiKeyAuth
// @Tags user
// @Description This request for get all users in company or department
// @ID getUsers
// @Accept json
// @Produces json
// @Param companyId query int false "company id"
// @Param departmentId query int false "department id"
// @Success 200 {array} core.User
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users [get]
func (H *Handler) getUsers(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! CurrentUser
// @Security ApiKeyAuth
// @Tags user
// @Description This request for edit current user info
// @ID patchCurrentUser
// @Accept json
// @Produces json
// @Param input body core.UserBaseData true "user info"
// @Success 200 {object} core.User
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [patch]
func (H *Handler) patchCurrentUser(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! CurrentUser
// @Security ApiKeyAuth
// @Tags user
// @Description This request for delete current user info
// @ID deleteCurrentUser
// @Accept json
// @Produces json
// @Success 200 {string} string "ok"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [delete]
func (H *Handler) deleteCurrentUser(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! LoadCurrentUserImage
// @Security ApiKeyAuth
// @Tags user
// @Description This request for load current user image
// @ID loadCurrentUserImage
// @Accept json
// @Produces json
// @Param input body []byte true "image file"
// @Success 200 {string} string "image_url"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/image [post]
func (H *Handler) loadCurrentUserImage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
