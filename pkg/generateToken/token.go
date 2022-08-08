package generateToken

import (
	"github.com/Thing-repository/backend-server/pkg/core"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	tokenTTL = time.Hour * 24 * 30
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId      int                `json:"user_id"`
	Credentials []core.Credentials `json:"credentials"`
}

type token struct {
	signingKey []byte
}

func NewToken(signingKey []byte) *token {
	return &token{signingKey: signingKey}
}

func (t *token) GenerateToken(userId int, credentials []core.Credentials) (string, error) {
	logBase := logrus.Fields{
		"module":   "generateToken",
		"function": "GenerateToken",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId:      userId,
		Credentials: credentials,
	})

	signetToken, err := token.SignedString(t.signingKey)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"userId": userId,
			"error":  err,
		}).Error("error generate token")
		return "", err
	}

	return signetToken, nil
}

func (t *token) ValidateToken(token string) (int, []core.Credentials, error) {
	logBase := logrus.Fields{
		"module":   "generateToken",
		"function": "ValidateToken",
	}

	res, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, moduleErrors.ErrorTokenInvalidToken
		}
		return t.signingKey, nil
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"token": token,
			"error": err,
		}).Warning("error validation token")
		return 0, nil, err
	}

	claims, ok := res.Claims.(*tokenClaims)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"token":  token,
			"claims": claims,
		}).Warning("error validation token")
		return 0, nil, moduleErrors.ErrorTokenInvalidToken
	}

	// validate date
	if claims.ExpiresAt < time.Now().Unix() {
		logrus.WithFields(logrus.Fields{
			"base":   logBase,
			"token":  token,
			"claims": claims,
		}).Warning("error token expired")
		return 0, nil, moduleErrors.ErrorTokenExpiredToken
	}

	return claims.UserId, claims.Credentials, nil
}
