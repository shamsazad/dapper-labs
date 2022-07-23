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

func Test_UserSignUp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		mockShop      func(mock *mocks.MockDaoInterface)
		apiCreateUser models.ApiCreateUser
		expectedError error
	}{
		{
			name: "Happy path, user able to signed up",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().CreateUser(gomock.Any()).Return(nil)
				mock.EXPECT().CreateHashedUserCredential(gomock.Any()).Return(nil)
			},
			apiCreateUser: fixtures.LoadCreateUserFixture("valid_user"),
			expectedError: nil,
		},
		{
			name: "invalid email address of User",
			mockShop: func(mock *mocks.MockDaoInterface) {
			},
			apiCreateUser: fixtures.LoadCreateUserFixture("invalid_email"),
			expectedError: errors.New("Email must be a valid email address"),
		},
		{
			name: "first name missing for user", //first name is required
			mockShop: func(mock *mocks.MockDaoInterface) {
			},
			apiCreateUser: fixtures.LoadCreateUserFixture("missing_first_name"),
			expectedError: errors.New("FirstName is a required field"),
		},
		{
			name: "password missing for user", //password is required
			mockShop: func(mock *mocks.MockDaoInterface) {
			},
			apiCreateUser: fixtures.LoadCreateUserFixture("missing_password"),
			expectedError: errors.New("Password is a required field"),
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
			err := service.UserSignUp(mockClient, test.apiCreateUser)
			if err != nil {
				assert.Equal(t, err, test.expectedError)
			}
		})
	}
}

func Test_GetAllUsers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		apiUsers      models.ApiUsers
		mockShop      func(mock *mocks.MockDaoInterface)
		expectedError error
	}{
		{
			name: "Happy path, users retrieved",
			apiUsers: models.ApiUsers{
				{
					Email:     "testUser@gmail.com",
					FirstName: "test",
					LastName:  "user",
				},
			},
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().GetAllUsers().Return(fixtures.LoadUserFixture("one_user"), nil)
			},
			expectedError: nil,
		},
		{
			name: "Error retriving users",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().GetAllUsers().Return(nil, errors.New("unable to connect to DB"))
			},
			expectedError: errors.New("unable to connect to DB"),
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
			apiUsers, err := service.GetAllUsers(mockClient)
			if err != nil {
				assert.Equal(t, err, test.expectedError)
			} else {
				assert.Equal(t, apiUsers, test.apiUsers)
			}
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		mockShop      func(mock *mocks.MockDaoInterface)
		expectedError error
	}{
		{
			name: "Happy path, user updated",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error updating user",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New("Out of transaction"))
			},
			expectedError: errors.New("Out of transaction"),
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
			err := service.UpdateUser(mockClient, models.ApiUpdateUser{}, "")
			if err != nil {
				assert.Equal(t, err, test.expectedError)
			}
		})
	}
}
