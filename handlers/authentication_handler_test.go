package handlers_test

import (
	"bytes"
	fixtures "dapper-labs/fixture"
	"dapper-labs/handlers"
	"dapper-labs/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Login(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		mockShop func(mock *mocks.MockDaoInterface)
		status   int
		payload  []byte
	}{
		{
			name: "bad request requested",
			mockShop: func(mock *mocks.MockDaoInterface) {
			},
			payload: []byte(""),
			status:  http.StatusInternalServerError,
		},
		{
			name: "happy path, users retrived",
			mockShop: func(mock *mocks.MockDaoInterface) {
				mock.EXPECT().FindHashedUserCredentials(gomock.Any()).Return(fixtures.LoadUserCredentialsFixture("valid_user_credential"), nil)
			},
			payload: []byte("{\n\t\t\t\t\n\t\t\t    \"email\":     \"test@test.com\",\n\t\t\t\t\"password\":  \"passwordoppo\"\n\t\t\t}"),
			status:  http.StatusCreated,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			body := bytes.NewReader(test.payload)
			req, err := http.NewRequest("POST", "/dapper-lab/login", body)
			if err != nil {
				t.Fatalf("Error creating a new request: %v", err)
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockClient := mocks.NewMockDaoInterface(ctrl)
			rr := httptest.NewRecorder()
			test.mockShop(mockClient)
			handler := http.HandlerFunc(handlers.Login(mockClient))
			handler.ServeHTTP(rr, req)

			assert.Equal(t, test.status, rr.Code)
		})
	}
}
