package models

import (
	"time"
)

type User struct {
	Email     string    `json:"email" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Users []User

func (u User) ConvertUserToApiUser() ApiCreateUser {
	return ApiCreateUser{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func (users Users) ConvertUsersToApiUsers() ApiUsers {
	var apiUsers []ApiCreateUser
	for _, user := range users {
		apiUser := user.ConvertUserToApiUser()
		apiUsers = append(apiUsers, apiUser)
	}
	return apiUsers
}
