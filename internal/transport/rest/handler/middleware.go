package handler

import (
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (H *Handler) userIdentity(c *gin.Context) {
	// TODO: realise check permission
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		logrus.WithFields(logrus.Fields{
			"module":         "handler",
			"function":       "userIdentity",
			"header":         header,
			"requestHeaders": c.Request.Header,
		}).Error("empty header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userId, credentials, err := H.token.ValidateToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}
	c.Set(core.UserIdCtx, userId)
	c.Set(core.CredentialsCtx, credentials)
}
