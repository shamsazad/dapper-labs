package service_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

type Suite struct {
	db   *gorm.DB
	mock sqlmock.Sqlmock
}
