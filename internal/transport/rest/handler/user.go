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

// @Summary NOT IMPLEMENTED! CurrentUser
// @Security ApiKeyAuth
// @Tags user
// @Description This request for edit current user info
// @ID patchCurrentUser
// @Accept json
// @Produces json
// @Param input body core.UserChange true "user info"
// @Success 200 {object} core.User
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [patch]
func (H *Handler) patchCurrentUser(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! User
// @Security ApiKeyAuth
// @Tags user
// @Description This request for edit user info
// @ID patchUser
// @Accept json
// @Produces json
// @Param id path int true "user id"
// @Param input body core.UserChange true "user info"
// @Success 200 {object} core.User
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id} [patch]
func (H *Handler) patchUser(c *gin.Context) {
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

// @Summary NOT IMPLEMENTED! User
// @Security ApiKeyAuth
// @Tags user
// @Description This request for delete user info
// @ID deleteUser
// @Accept json
// @Produces json
// @Param id path int true "user id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id} [delete]
func (H *Handler) deleteUser(c *gin.Context) {
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

// @Summary NOT IMPLEMENTED! LoadUserImage
// @Security ApiKeyAuth
// @Tags user
// @Description This request for load user image
// @ID loadUserImage
// @Accept json
// @Produces json
// @Param id path int true "user id"
// @Param input body []byte true "image file"
// @Success 200 {string} string "image_url"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/{id}/image [post]
func (H *Handler) loadUserImage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
