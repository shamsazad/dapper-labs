package handlers_test

import (
	"bytes"
	fixtures "dapper-labs/fixture"
	"dapper-labs/handlers"
	"dapper-labs/mocks"
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetAllUsers(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		mockShop func(mock *mocks.MockDaoInterface)
		status   int
	}{
		{
			name: "bad request requested",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().GetAllUsers().Return(nil, errors.New("unable to reach database"))
			},
			status: http.StatusBadRequest,
		},
		{
			name: "happy path, users retrived",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().GetAllUsers().Return(fixtures.LoadUserFixture("two_users"), nil)
			},
			status: http.StatusOK,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequest("GET", "/auth/dapper-lab/users", errReader(0))
			if err != nil {
				t.Fatalf("Error creating a new request: %v", err)
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockDaoInterface(ctrl)
			rr := httptest.NewRecorder()
			test.mockShop(mockClient)
			handler := http.HandlerFunc(handlers.GetAllUsers(mockClient))
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.status, rr.Code)
		})
	}
}

func Test_SignUp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		mockShop func(mock *mocks.MockDaoInterface)
		status   int
		payload  []byte
	}{
		{
			name: "empty create user body",
			mockShop: func(mock *mocks.MockDaoInterface) {
			},
			payload: []byte(""),
			status:  http.StatusBadRequest,
		},
		{
			name: "happy path, users able to signUp",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().CreateUser(gomock.Any()).Return(nil)
				mock.EXPECT().CreateHashedUserCredential(gomock.Any()).Return(nil)
			},
			payload: []byte("{\n\t\t\t\t\"first_name\": \"test\",\n\t\t\t\t\"last_name\":  \"user\",\n\t\t\t\t\"email\":     \"test@test.com\",\n\t\t\t\t\"password\":  \"password\"\n\t\t\t}"),
			status:  http.StatusCreated,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			body := bytes.NewReader(test.payload)
			req, err := http.NewRequest("POST", "/dapper-lab/user", body)
			if err != nil {
				t.Fatalf("Error creating a new request: %v", err)
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockDaoInterface(ctrl)
			rr := httptest.NewRecorder()
			test.mockShop(mockClient)
			handler := http.HandlerFunc(handlers.SignUp(mockClient))
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.status, rr.Code)
		})
	}
}

func Test_Update(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		mockShop func(mock *mocks.MockDaoInterface)
		status   int
		payload  []byte
	}{
		{
			name: "update empty user",
			mockShop: func(mock *mocks.MockDaoInterface) {
			},
			payload: []byte(""),
			status:  http.StatusBadRequest,
		},
		{
			name: "happy path, users able to update",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)
			},
			payload: []byte("{\n\t\t\t\t\"first_name\": \"test\",\n\t\t\t\t\"last_name\":  \"user\"}"),
			status:  http.StatusCreated,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			body := bytes.NewReader(test.payload)
			req, err := http.NewRequest("POST", "auth/dapper-lab/update/user", body)
			if err != nil {
				t.Fatalf("Error creating a new request: %v", err)
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockDaoInterface(ctrl)
			rr := httptest.NewRecorder()
			test.mockShop(mockClient)
			handler := http.HandlerFunc(handlers.UpdateUser(mockClient))
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.status, rr.Code)
		})
	}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}
