package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary NOT IMPLEMENTED! Thing usage
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for get all usage info in company
// @ID getAllCompanyThingUsage
// @Accept json
// @Produces json
// @Param companyId query int false "company id"
// @Success 200 {array} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/company [get]
func (H *Handler) getAllCompanyThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for get all usage info in department
// @ID getAllDepartmentThingUsage
// @Accept json
// @Produces json
// @Param companyId query int false "company id"
// @Param departmentId query int false "department id"
// @Success 200 {array} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/department [get]
func (H *Handler) getAllDepartmentThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for using thing
// @ID addThingUsage
// @Accept json
// @Produces json
// @Param id query int true "thing id"
// @Param input body core.ThingUsageAdd true "thing use info"
// @Success 200 {object} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage [post]
func (H *Handler) addThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for edit usage info
// @ID editThingUsage
// @Accept json
// @Produces json
// @Param id path int true "usage id"
// @Param input body core.ThingUsageAdd true "thing usage info"
// @Success 200 {object} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/{id} [patch]
func (H *Handler) editThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for get usage info
// @ID getThingUsage
// @Accept json
// @Produces json
// @Param id path int true "usage id"
// @Success 200 {object} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/{id} [get]
func (H *Handler) getThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for delete usage
// @ID deleteThingUsage
// @Accept json
// @Produces json
// @Param id path int true "usage id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/{id} [delete]
func (H *Handler) deleteThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage get need approve
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for get all things need approved by me
// @ID getNeedApproveThingUsage
// @Accept json
// @Produces json
// @Success 200 {array} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/approve [get]
func (H *Handler) getNeedApproveThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage approve
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for approve usage
// @ID approveThingUsage
// @Accept json
// @Produces json
// @Param id path int true "usage id"
// @Success 200 {object} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/{id}/approve [post]
func (H *Handler) approveThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Get my usage
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for get all things usage by me
// @ID getMyThingUsage
// @Accept json
// @Produces json
// @Success 200 {array} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/my [get]
func (H *Handler) getMyThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage take
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for take thing
// @ID takeThingUsage
// @Accept json
// @Produces json
// @Param id path int true "usage id"
// @Success 200 {object} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/{id}/take [post]
func (H *Handler) takeThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing usage return
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for return thing
// @ID returnThingUsage
// @Accept json
// @Produces json
// @Param id path int true "usage id"
// @Success 200 {object} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/usage/{id}/return [post]
func (H *Handler) returnThingUsage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Get usage by thing
// @Security ApiKeyAuth
// @Tags thing usage
// @Description This request for get all usage by thing
// @ID getThingUsageById
// @Accept json
// @Produces json
// @Success 200 {array} core.ThingUsage
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/{id}/usage/ [get]
func (H *Handler) getThingUsageById(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
