package service_test

import (
	fixtures "dapper-labs/fixture"
	"dapper-labs/mocks"
	"dapper-labs/models"
	"dapper-labs/service"
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_Login(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		mockShop        func(mock *mocks.MockDaoInterface)
		loginCredential models.LoginCredential
		expectedError   error
	}{
		{
			name: "Happy path, user able to Login",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().FindHashedUserCredentials(gomock.Any()).Return(fixtures.LoadUserCredentialsFixture("valid_user_credential"), nil)
			},
			loginCredential: models.LoginCredential{
				Email:    "shams@gmail.com",
				Password: "passwordoppo",
			},
			expectedError: nil,
		},
		{
			name: "login password is not same as hashed password",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().FindHashedUserCredentials(gomock.Any()).Return(fixtures.LoadUserCredentialsFixture("invalid_user_credential"), nil)
			},
			loginCredential: models.LoginCredential{
				Email:    "shams@gmail.com",
				Password: "password",
			},
			expectedError: errors.New("crypto/bcrypt: hashedPassword is not the hash of the given password"),
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockDaoInterface(ctrl)
			test.mockShop(mockClient)
			err := service.Login(mockClient, test.loginCredential)
			if err != nil {
				assert.Equal(t, err, test.expectedError)
			}
		})
	}
}
