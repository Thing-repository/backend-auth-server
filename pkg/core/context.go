package core

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
)

const (
	UserDataCtx    = "userData"
	UserIdCtx      = "userId"
	CredentialsCtx = "credentials"
)

var allContextValues = []string{UserDataCtx, UserIdCtx, CredentialsCtx}

func ContextGetUserId(ctx context.Context) (int, error) {
	userId := ctx.Value(UserIdCtx)
	if userId == nil {
		return 0, errors.New("can't get user id")
	}
	id, ok := userId.(int)
	if !ok {
		return 0, errors.New("invalid user id")
	}
	return id, nil
}

func ContextGetUserCredentials(ctx context.Context) ([]Credentials, error) {
	userCredentials := ctx.Value(CredentialsCtx)
	if userCredentials == nil {
		return nil, errors.New("can't get user credentials")
	}
	credentials, ok := userCredentials.([]Credentials)
	if !ok {
		return nil, errors.New("invalid user credentials")
	}
	return credentials, nil
}

func LogContext(ctx context.Context) *logrus.Fields {
	fields := make(logrus.Fields)
	for _, val := range allContextValues {
		fields[val] = ctx.Value(val)
	}
	return &fields
}
