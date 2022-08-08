package handler

import (
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// @Summary company
// @Security ApiKeyAuth
// @Tags company
// @Description This request for creating company
// @ID addCompany
// @Accept json
// @Produces json
// @Param input body core.CompanyBase true "company info"
// @Success 200 {object} core.Company
// @Failure 400,401,409 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company [post]
func (H *Handler) addCompany(c *gin.Context) {
	logBase := logrus.Fields{
		"module":   "handler",
		"function": "addCompany",
		"context":  *core.LogContext(c),
	}

	var company core.CompanyBase

	if err := c.BindJSON(&company); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("json parse error")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if company.CompanyName == nil || company.Address == nil {
		logrus.WithFields(logrus.Fields{
			"base":    logBase,
			"company": company,
		}).Error(moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
	}

	res, err := H.company.AddCompany(c, &company)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("add company error")
		switch err {
		case moduleErrors.ErrorServiceUserAlreadyHasCompany:
			newErrorResponse(c, http.StatusConflict, moduleErrors.ErrorServiceUserAlreadyHasCompany.Error())
			return
		case moduleErrors.ErrorServiceInvalidContext:
			newErrorResponse(c, http.StatusUnauthorized, moduleErrors.ErrorHandlerForbidden.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, res)
}

// @Summary company
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
	logBase := logrus.Fields{
		"module":   "handler",
		"function": "getCompany",
	}

	companyId, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error(moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
	}

	companyData, err := H.company.GetCompany(c, companyId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("get company error")
		switch err {
		case moduleErrors.ErrorServiceInvalidContext:
			newErrorResponse(c, http.StatusUnauthorized, moduleErrors.ErrorHandlerForbidden.Error())
			return
		case moduleErrors.ErrorServiceBadPermissions:
			newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorServiceBadPermissions.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, companyData)
}

// @Summary company
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
	logBase := logrus.Fields{
		"module":   "handler",
		"function": "patchCompany",
		"context":  *core.LogContext(c),
	}

	companyId, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error(moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
	}

	var company core.CompanyBase

	if err := c.BindJSON(&company); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("json parse error")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !(company.CompanyName != nil || company.Address != nil) {
		logrus.WithFields(logrus.Fields{
			"base":    logBase,
			"company": company,
			"error":   err.Error(),
		}).Error("nothing to change")
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorAllNoFields.Error())
		return
	}

	companyData, err := H.company.UpdateCompany(c, company, companyId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"company":   company,
			"companyId": companyId,
			"error":     err.Error(),
		}).Error("update company error")
		switch err {
		case moduleErrors.ErrorServiceInvalidContext:
			newErrorResponse(c, http.StatusUnauthorized, moduleErrors.ErrorHandlerForbidden.Error())
			return
		case moduleErrors.ErrorServiceBadPermissions:
			newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorServiceBadPermissions.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, companyData)
}

// @Summary company
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
	logBase := logrus.Fields{
		"module":   "handler",
		"function": "deleteCompany",
	}

	companyId, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error(moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorHandlerNoRequiredFieldsQuery.Error())
	}

	err = H.company.DeleteCompany(c, companyId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
			"error":     err.Error(),
		}).Error("delete company error")
		switch err {
		case moduleErrors.ErrorServiceInvalidContext:
			newErrorResponse(c, http.StatusUnauthorized, moduleErrors.ErrorHandlerForbidden.Error())
			return
		case moduleErrors.ErrorServiceBadPermissions:
			newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorServiceBadPermissions.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, "ok")
}

// @Summary NOT IMPLEMENTED! LoadCompanyImage
// @Security ApiKeyAuth
// @Tags company
// @Description This request for load company image
// @ID loadCompanyImage
// @Accept json
// @Produces json
// @Param id path int true "company id"
// @Param input body []byte true "image file"
// @Success 200 {string} string "image_url"
// @Failure 400,401 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /company/{id}/image [post]
func (H *Handler) loadCompanyImage(c *gin.Context) {
	newErrorResponse(c, http.StatusInternalServerError, "method not implemented")
}
