package fixtures

import (
	"dapper-labs/models"
	"github.com/mohae/deepcopy"
	"log"
)

var createUserCredentialFixture map[string]models.UserCredential

func init() {
	createUserCredentialFixture = make(map[string]models.UserCredential)
	createUserCredentialFixture = map[string]models.UserCredential{
		"valid_user_credential": {
			HashedPassword: "$2a$10$wkBf3g0NO6Eu2g2zMi1sweroBHdOVNll77xX4u74Yi463SnNllTiG",
			Email:          "test@test.com",
		},
		"invalid_user_credential": {
			HashedPassword: "$2a$10$IEycUS.P6BYa2fB.vnkqBu0GcglZhYlMPdKS6ojKRxp0pLq9ovYhu",
			Email:          "test@test.com",
		},
	}
}

func LoadUserCredentialsFixture(name string) models.UserCredential {
	fixture, ok := createUserCredentialFixture[name]
	if !ok {
		log.Fatalf("No fixture of type %T with name '%v' found", fixture, name)
	}
	newUserCredential := deepcopy.Copy(fixture).(models.UserCredential)

	return newUserCredential
}
