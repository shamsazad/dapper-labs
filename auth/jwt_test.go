package auth_test

import (
	"dapper-labs/auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ValidateJWT(t *testing.T) {

	t.Parallel()
	testCases := []struct {
		name     string
		token    string
		username string
	}{
		{
			name:     "Happy path, JWT Validated",
			token:    createValidToken("shams@gmail.com"),
			username: "shams@gmail.com",
		},
		{
			name:     "JWT not validated",
			token:    "invalid token",
			username: "shams@gmail.com",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			userName, err := auth.ValidateToken(test.token)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.username, userName)
			}
		})
	}

}

func Test_GenerateJWT(t *testing.T) {

	t.Parallel()
	testCases := []struct {
		name          string
		email         string
		expectedError error
	}{
		{
			name:  "Happy path, JWT created",
			email: "shams@gmail.com",
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			token, err := auth.GenerateJWT(test.email)
			assert.NotNil(t, token)
			assert.Nil(t, err)
		})
	}

}

func createValidToken(email string) string {
	token, _ := auth.GenerateJWT(email)
	return token
}
