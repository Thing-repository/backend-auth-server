package handler

import (
	"context"
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
	}
	userId, err := getUserId(c)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user id")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := context.TODO()

	userData, err := H.userDB.GetUser(ctx, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user data")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userData.CompanyId != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": userData,
		}).Error(moduleErrors.ErrorHandlerUserAlreadyHasCompany.Error())
		newErrorResponse(c, http.StatusConflict, moduleErrors.ErrorHandlerUserAlreadyHasCompany.Error())
		return
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

	res, err := H.company.AddCompany(&company, &userData.User)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("add company error")
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

	userId, err := getUserId(c)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user id")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := context.TODO()

	companyId, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error(moduleErrors.ErrorNoRequiredFieldsQuery.Error())
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorNoRequiredFieldsQuery.Error())
	}

	userData, err := H.userDB.GetUser(ctx, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user data")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: add check rights for service admin
	if userData.CompanyId == nil || *userData.CompanyId != companyId {
		if userData.CompanyId == nil {
			logrus.WithFields(logrus.Fields{
				"base": logBase,
			}).Error("user has not companyId")
		} else {
			logrus.WithFields(logrus.Fields{
				"base":            logBase,
				"userCompany":     *userData.CompanyId,
				"requiredCompany": companyId,
			}).Error("user has invalid companyId")
		}

		newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorForbidden.Error())
		return
	}

	companyData, err := H.company.GetCompany(companyId)

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
	}

	userId, err := getUserId(c)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user id")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := context.TODO()

	companyId, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error(moduleErrors.ErrorNoRequiredFieldsQuery.Error())
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorNoRequiredFieldsQuery.Error())
	}

	userData, err := H.userDB.GetUser(ctx, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user data")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: add check rights for service admin
	if true {
		if userData.CompanyId == nil || *userData.CompanyId != companyId {
			if userData.CompanyId == nil {
				logrus.WithFields(logrus.Fields{
					"base": logBase,
				}).Error("user has not companyId")
			} else {
				logrus.WithFields(logrus.Fields{
					"base":            logBase,
					"userCompany":     *userData.CompanyId,
					"requiredCompany": companyId,
				}).Error("user has invalid companyId")
			}

			newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorForbidden.Error())
			return
		}
		// TODO: realise check permission
		if userData.CompanyId == nil || *userData.CompanyId != companyId {
			if userData.CompanyId == nil {
				logrus.WithFields(logrus.Fields{
					"base": logBase,
				}).Error("user has not companyId")
			} else {
				logrus.WithFields(logrus.Fields{
					"base":            logBase,
					"userCompany":     *userData.CompanyId,
					"requiredCompany": companyId,
				}).Error("user has invalid companyId")
			}

			newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorForbidden.Error())
			return
		}
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

	companyData, err := H.company.UpdateCompany(company, companyId)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"company":   company,
			"companyId": companyId,
			"error":     err.Error(),
		}).Error("update company error")
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

	userId, err := getUserId(c)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user id")
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := context.TODO()

	companyId, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error(moduleErrors.ErrorNoRequiredFieldsQuery.Error())
		newErrorResponse(c, http.StatusBadRequest, moduleErrors.ErrorNoRequiredFieldsQuery.Error())
	}

	userData, err := H.userDB.GetUser(ctx, userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("error get user data")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: add check rights for service admin
	if true {
		if userData.CompanyId == nil || *userData.CompanyId != companyId {
			if userData.CompanyId == nil {
				logrus.WithFields(logrus.Fields{
					"base": logBase,
				}).Error("user has not companyId")
			} else {
				logrus.WithFields(logrus.Fields{
					"base":            logBase,
					"userCompany":     *userData.CompanyId,
					"requiredCompany": companyId,
				}).Error("user has invalid companyId")
			}

			newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorForbidden.Error())
			return
		}
		// TODO: realise check permission
		if userData.CompanyId == nil || *userData.CompanyId != companyId {
			if userData.CompanyId == nil {
				logrus.WithFields(logrus.Fields{
					"base": logBase,
				}).Error("user has not companyId")
			} else {
				logrus.WithFields(logrus.Fields{
					"base":            logBase,
					"userCompany":     *userData.CompanyId,
					"requiredCompany": companyId,
				}).Error("user has invalid companyId")
			}

			newErrorResponse(c, http.StatusForbidden, moduleErrors.ErrorForbidden.Error())
			return
		}
	}

	err = H.company.DeleteCompany(companyId)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":      logBase,
			"companyId": companyId,
			"error":     err.Error(),
		}).Error("delete company error")
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
