package service

import (
	"github.com/Thing-repository/backend-server/pkg/core"
)

func newCredential(objectId int, userId int, credentialType string) *core.AddCredentials {
	return &core.AddCredentials{
		UserId: userId,
		Credentials: core.Credentials{
			CredentialType: credentialType,
			ObjectId:       objectId,
		},
	}
}
