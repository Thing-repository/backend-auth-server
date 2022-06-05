package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary NOT IMPLEMENTED! CompanyAdmins
// @Security ApiKeyAuth
// @Tags rights
// @Description This request for add user to company admins
// @ID addCompanyAdmin
// @Accept json
// @Produces json
// @Param userId query int true "added user id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company_admins [post]
func (H *Handler) addCompanyAdmin(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! CompanyAdmins
// @Security ApiKeyAuth
// @Tags rights
// @Description This request for delete user from company admins
// @ID deleteCompanyAdmin
// @Accept json
// @Produces json
// @Param userId query int true "deleted user id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company_admins [delete]
func (H *Handler) deleteCompanyAdmin(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! DepartmentAdmin
// @Security ApiKeyAuth
// @Tags rights
// @Description This request for add user to department admins
// @ID addDepartmentAdmin
// @Accept json
// @Produces json
// @Param userId query int true "added user id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department_admins [post]
func (H *Handler) addDepartmentAdmin(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! DepartmentAdmin
// @Security ApiKeyAuth
// @Tags rights
// @Description This request for delete user from department admins
// @ID deleteDepartmentAdmin
// @Accept json
// @Produces json
// @Param userId query int true "deleted user id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department_admins [delete]
func (H *Handler) deleteDepartmentAdmin(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! DepartmentMaintainer
// @Security ApiKeyAuth
// @Tags rights
// @Description This request for add user to department maintainers
// @ID addDepartmentMaintainer
// @Accept json
// @Produces json
// @Param userId query int true "added user id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department_maintainers [post]
func (H *Handler) addDepartmentMaintainer(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! DepartmentMaintainer
// @Security ApiKeyAuth
// @Tags rights
// @Description This request for delete user from department maintainers
// @ID deleteDepartmentMaintainer
// @Accept json
// @Produces json
// @Param userId query int true "deleted user id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /department_maintainers [delete]
func (H *Handler) deleteDepartmentMaintainer(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
