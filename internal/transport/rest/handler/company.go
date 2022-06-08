package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary NOT IMPLEMENTED! Company
// @Security ApiKeyAuth
// @Tags company
// @Description This request for creating company
// @ID addCompany
// @Accept json
// @Produces json
// @Param input body core.CompanyBase true "company info"
// @Success 200 {object} core.Company
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company [post]
func (H *Handler) addCompany(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Company
// @Security ApiKeyAuth
// @Tags company
// @Description This request for get company info
// @ID getCompany
// @Accept json
// @Produces json
// @Param id path int true "company id"
// @Success 200 {object} core.Company
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company/{id} [get]
func (H *Handler) getCompany(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Company
// @Security ApiKeyAuth
// @Tags company
// @Description This request for change company info
// @ID patchCompany
// @Accept json
// @Produces json
// @Param id path int true "company id"
// @Param input body core.CompanyBase true "new company info"
// @Success 200 {object} core.Company
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company/{id} [patch]
func (H *Handler) patchCompany(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! Company
// @Security ApiKeyAuth
// @Tags company
// @Description This request for delete company
// @ID deleteCompany
// @Accept json
// @Produces json
// @Param id path int true "company id"
// @Success 200 {string} string "ok"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company/{id} [delete]
func (H *Handler) deleteCompany(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}

// @Summary NOT IMPLEMENTED! LoadCompanyImage
// @Security ApiKeyAuth
// @Tags company
// @Description This request for load company image
// @ID loadCompanyImage
// @Accept json
// @Produces json
// @Param input body []byte true "image file"
// @Success 200 {string} string "image_url"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company/{id}/image [post]
func (H *Handler) loadCompanyImage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
