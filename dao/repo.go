package dao

import "gorm.io/gorm"

type Repo struct {
	DB *gorm.DB
}
