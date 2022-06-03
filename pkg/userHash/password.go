package userHash

import (
	"crypto/sha1"
	"fmt"
	"github.com/Thing-repository/backend-server/pkg/core/moduleErrors"
	"github.com/sirupsen/logrus"
)

type Hash struct {
	salt string
}

func NewHash(salt string) *Hash {
	return &Hash{salt: salt}
}

func (h *Hash) GenerateHash(password string) (string, error) {
	logBase := logrus.Fields{
		"module":   "userHash",
		"function": "GenerateHash",
	}
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		logrus.WithFields(logrus.Fields{
			"base":  logBase,
			"error": err,
		}).Error("error generate hash")
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}

func (h *Hash) ValidateHash(hash string, password string) error {
	logBase := logrus.Fields{
		"module":   "userHash",
		"function": "ValidateHash",
	}
	calculatedHash, err := h.GenerateHash(password)
	if err != nil {
		return err
	}
	if hash != calculatedHash {
		logrus.WithFields(logrus.Fields{
			"base":           logBase,
			"calculatedHash": calculatedHash,
			"hash":           hash,
		}).Warning("hashes doesn't match'")
		return moduleErrors.ErrorHashValidationPassword
	}
	return nil
}
