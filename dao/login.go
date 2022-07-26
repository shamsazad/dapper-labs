package dao

import (
	"dapper-labs/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginInterface interface {
	CreateHashedUserCredential(models.LoginCredential) error
	FindHashedUserCredentials(email string) (hashedCredential models.UserCredential, err error)
}

func (d *Dao) CreateHashedUserCredential(credential models.LoginCredential) error {

	var hashedCredential models.UserCredential
	bytes, err := bcrypt.GenerateFromPassword([]byte(credential.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedCredential.HashedPassword = string(bytes)
	hashedCredential.Email = credential.Email
	if err = d.DB.Create(&hashedCredential).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) FindHashedUserCredentials(email string) (hashedCredential models.UserCredential, err error) {

	if err = d.DB.First(&hashedCredential, "email = ?", email).Error; err != nil {
		return hashedCredential, err
	}
	return hashedCredential, nil
}
