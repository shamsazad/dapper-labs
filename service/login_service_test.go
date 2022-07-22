package service_test

import (
	"dapper-labs/models"
	//"dapper-labs/service"
	//"fmt"
	//"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
	"testing"
)

func Test_UserSingUp(t *testing.T) {
	t.Parallel()
	//db, s.mock, err := sqlmock.New()
	//if err != nil {
	//	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	//}
	//defer db.Close()

	testCases := []struct {
		name    string
		apiUser models.ApiCreateUser
		db      *gorm.DB
	}{
		{
			name: "just test",
			apiUser: models.ApiCreateUser{
				Email:     "shams@gmail.com",
				FirstName: "Babu",
				LastName:  "bhai",
				Password:  "password",
			},
			//db: db, mock, err := sqlmock.New()
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			//err := service.UserSignUp(test.db, test.apiUser)
			//if err != nil {
			//	fmt.Println(err)
			//}
		})
	}
}
