package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary NOT IMPLEMENTED! Department
// @Security ApiKeyAuth
// @Tags department
// @Description This request for creating Department
// @ID addDepartment
// @Accept json
// @Produces json
// @Param input body core.DepartmentBase true "department info"
// @Success 200 {object} core.Department
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department [post]
func (H *Handler) addDepartment(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Department
// @Security ApiKeyAuth
// @Tags department
// @Description This request for get Department info
// @ID getDepartment
// @Accept json
// @Produces json
// @Param id path int true "department id"
// @Success 200 {object} core.Department
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department/{id} [get]
func (H *Handler) getDepartment(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Department
// @Security ApiKeyAuth
// @Tags department
// @Description This request for change department info
// @ID patchDepartment
// @Accept json
// @Produces json
// @Param id path int true "department id"
// @Param input body core.DepartmentBase true "new department info"
// @Success 200 {object} core.Department
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department/{id} [patch]
func (H *Handler) patchDepartment(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Department
// @Security ApiKeyAuth
// @Tags department
// @Description This request for delete department
// @ID deleteDepartment
// @Accept json
// @Produces json
// @Param id path int true "department id"
// @Success 200 {string} string "ok"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department/{id} [delete]
func (H *Handler) deleteDepartment(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! LoadDepartmentImage
// @Security ApiKeyAuth
// @Tags department
// @Description This request for load department image
// @ID loadDepartmentImage
// @Accept json
// @Produces json
// @Param input body []byte true "image file"
// @Success 200 {string} string "image_url"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department/{id}/image [post]
func (H *Handler) loadDepartmentImage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
