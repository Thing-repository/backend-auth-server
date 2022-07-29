package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary NOT IMPLEMENTED! Thing
// @Security ApiKeyAuth
// @Tags thing
// @Description This request for creating thing
// @ID addThing
// @Accept json
// @Produces json
// @Param input body core.ThingBase true "thing info"
// @Param departmentId query int false "department id"
// @Success 200 {object} core.Thing
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing [post]
func (H *Handler) addThing(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing
// @Security ApiKeyAuth
// @Tags thing
// @Description This request for getting thing
// @ID getThing
// @Accept json
// @Produces json
// @Param id path int true "thing id"
// @Success 200 {object} core.Thing
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/{id} [get]
func (H *Handler) getThing(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! All things
// @Security ApiKeyAuth
// @Tags thing
// @Description This request for getting all company or department things
// @ID getAllThings
// @Accept json
// @Produces json
// @Param companyId query int false "company id"
// @Param departmentId query int false "department id"
// @Success 200 {array} core.Thing
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /things [get]
func (H *Handler) getAllThings(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing
// @Security ApiKeyAuth
// @Tags thing
// @Description This request for edit thing info
// @ID patchThing
// @Accept json
// @Produces json
// @Param id path int true "thing id"
// @Param input body core.ThingBase true "thing info"
// @Success 200 {object} core.Thing
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/{id} [patch]
func (H *Handler) patchThing(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing
// @Security ApiKeyAuth
// @Tags thing
// @Description This request for delete thing
// @ID deleteThing
// @Accept json
// @Produces json
// @Param id path int true "thing id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/{id} [delete]
func (H *Handler) deleteThing(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
