package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Router /user/{id} [get]
func (H *Handler) getUser(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
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
