package dao

import (
	"dapper-labs/models"
)

type UserInterface interface {
	CreateUser(models.ApiCreateUser) error
	UpdateUser(models.ApiUpdateUser, string) error
	GetAllUsers() (apiUsers models.ApiUsers, err error)
}

func (d *Dao) CreateUser(apiUser models.ApiCreateUser) error {

	user := apiUser.ConvertApiUserToUser()
	if err := d.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) UpdateUser(apiUpdateUser models.ApiUpdateUser, email string) error {

	var user models.User
	if err := d.DB.First(&user, "email = ?", email).Error; err != nil {
		return err
	}
	if err := d.DB.Model(&user).Updates(models.User{FirstName: apiUpdateUser.FirstName, LastName: apiUpdateUser.LastName}).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetAllUsers() (apiUsers models.ApiUsers, err error) {
	var users models.Users
	if err = d.DB.Find(&users).Error; err != nil {
		return apiUsers, err
	}
	apiUsers = users.ConvertUsersToApiUsers()
	return apiUsers, nil
}
