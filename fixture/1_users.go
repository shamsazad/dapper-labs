package fixtures

import (
	"dapper-labs/models"
	"github.com/mohae/deepcopy"
	"log"
)

var userFixtures map[string]models.ApiUsers
var createUserFixture map[string]models.ApiCreateUser

func init() {
	userFixtures = make(map[string]models.ApiUsers)
	createUserFixture = make(map[string]models.ApiCreateUser)

	userFixtures = map[string]models.ApiUsers{
		"one_user": {
			{
				Email:     "testUser@gmail.com",
				FirstName: "test",
				LastName:  "user",
			},
		},
		"two_users": {
			{
				Email:     "testUser2@gmail.com",
				FirstName: "test2",
				LastName:  "user2",
			},
			{
				Email:     "testUser3@gmail.com",
				FirstName: "test3",
				LastName:  "user3",
			},
		},
	}
	createUserFixture = map[string]models.ApiCreateUser{
		"valid_user": {
			FirstName: "test",
			LastName:  "user",
			Email:     "test@test.com",
			Password:  "password",
		},
		"invalid_email": {
			FirstName: "test",
			LastName:  "user",
			Email:     "test@test", //invalid email
			Password:  "password",
		},
		"missing_first_name": { //first name is required
			LastName: "user",
			Email:    "test@test.com",
			Password: "password",
		},
		"missing_password": { //password is required
			FirstName: "test",
			LastName:  "user",
			Email:     "test@test.com",
		},
	}
}

func LoadUserFixture(name string) models.ApiUsers {
	fixture, ok := userFixtures[name]
	if !ok {
		log.Fatalf("No fixture of type %T with name '%v' found", fixture, name)
	}
	newUsers := deepcopy.Copy(fixture).(models.ApiUsers)

	return newUsers
}

func LoadCreateUserFixture(name string) models.ApiCreateUser {
	fixture, ok := createUserFixture[name]
	if !ok {
		log.Fatalf("No fixture of type %T with name '%v' found", fixture, name)
	}
	newCreateUsers := deepcopy.Copy(fixture).(models.ApiCreateUser)

	return newCreateUsers
}
