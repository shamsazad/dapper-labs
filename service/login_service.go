package service

import (
	"dapper-labs/dao"
	"dapper-labs/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(DAO dao.LoginInterface, loginCredential models.LoginCredential) error {

	hashedCredential, err := DAO.FindHashedUserCredentials(loginCredential.Email)
	if err != nil {
		return err
	}

	err = checkPassword(hashedCredential.HashedPassword, loginCredential.Password)
	if err != nil {
		return err
	}

	return nil
}

func checkPassword(hashedPassword string, loginPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginPassword))
	if err != nil {
		return err
	}
	return nil
}
