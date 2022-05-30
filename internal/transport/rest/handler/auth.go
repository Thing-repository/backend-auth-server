package handler

import (
	"errors"
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

const (
	MinPasswordLength = 8
)

func validateEmail(email string) error {
	regExpEmail := regexp.MustCompile("^[^\\[\\]\\\\;,\\s]*@[^\\[\\]\\\\;,\\s]*\\.[^\\[\\]\\\\;,\\s]*")
	if !regExpEmail.MatchString(email) {
		return errors.New("invalid email")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < MinPasswordLength {
		return errors.New("invalid password, too short")
	}
	if !regexp.MustCompile(".*\\d").MatchString(password) {
		return errors.New("invalid password, no numbers")
	}
	if !regexp.MustCompile(".*[A-Z]").MatchString(password) {
		return errors.New("invalid password, no uppercase letters")
	}
	if !regexp.MustCompile(".*[a-z]").MatchString(password) {
		return errors.New("invalid password, no lowercase letters")
	}

	return nil
}

func (H *Handler) signUp(c *gin.Context) {
	logBase := logrus.Fields{
		"module":   "handler",
		"function": "signUp",
	}

	var input core.UserSignUpData

	if err := c.BindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("json parsing error")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateEmail(input.Email); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
			"email": input.Email,
		}).Error("invalid email address")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validatePassword(input.Password); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("invalid password")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userData, err := H.auth.SignUp(&input)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err.Error(),
		}).Error("sgn-up error")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userData)

}

func (H *Handler) signIn(c *gin.Context) {
	logBase := logrus.Fields{
		"module":   "handler",
		"function": "signIn",
	}

	var input core.UserSignInData

	if err := c.BindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("json parsing error")
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.UserMail == "" || input.UserPassword == "" || len(input.UserPassword) < MinPasswordLength {
		logrus.WithFields(logrus.Fields{
			"base": logBase,
		}).Error("Bad username or password")
		newErrorResponse(c, http.StatusBadRequest, "Invalid username or password")
		return
	}

	userData, err := H.auth.SignIn(&input)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("authorization error")
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userData)
}
