package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary NOT IMPLEMENTED! Thing block
// @Security ApiKeyAuth
// @Tags thing block
// @Description This request for get all block info in company
// @ID getAllCompanyThingBlocking
// @Accept json
// @Produces json
// @Param companyId query int false "company id"
// @Success 200 {array} core.ThingBlock
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/block/company [get]
func (H *Handler) getAllCompanyThingBlocking(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing block
// @Security ApiKeyAuth
// @Tags thing block
// @Description This request for get all block info in department
// @ID getAllDepartmentThingBlocking
// @Accept json
// @Produces json
// @Param companyId query int false "company id"
// @Param departmentId query int false "department id"
// @Success 200 {array} core.ThingBlock
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/block/department [get]
func (H *Handler) getAllDepartmentThingBlocking(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing block
// @Security ApiKeyAuth
// @Tags thing block
// @Description This request for blocking thing
// @ID addThingBlock
// @Accept json
// @Produces json
// @Param id query int true "thing id"
// @Param input body core.ThingBlockAdd true "thing bock info"
// @Success 200 {object} core.ThingBlock
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/block [post]
func (H *Handler) addThingBlock(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing block
// @Security ApiKeyAuth
// @Tags thing block
// @Description This request for edit block info
// @ID editThingBlock
// @Accept json
// @Produces json
// @Param id path int true "block id"
// @Param input body core.ThingBlockAdd true "thing bock info"
// @Success 200 {object} core.ThingBlock
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/block/{id} [patch]
func (H *Handler) editThingBlock(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing block
// @Security ApiKeyAuth
// @Tags thing block
// @Description This request for get block info
// @ID getThingBlock
// @Accept json
// @Produces json
// @Param id path int true "block id"
// @Success 200 {object} core.ThingBlock
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/block/{id} [get]
func (H *Handler) getThingBlocking(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Thing block
// @Security ApiKeyAuth
// @Tags thing block
// @Description This request for delete block
// @ID addThingBlocking
// @Accept json
// @Produces json
// @Param id path int true "block id"
// @Success 200 {string} string "ok"
// @Failure 400,401,403 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /thing/block/{id} [delete]
func (H *Handler) deleteThingBlocking(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
