package service

import (
	"dapper-labs/dao"
	"dapper-labs/models"
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

//Signup need to run in transaction
func UserSignUp(repo dao.Repo, apiUser models.ApiCreateUser) error {

	errs := validateUser(apiUser)
	if errs != "" {
		return errors.New(errs)
	}

	if err := repo.CreateUser(apiUser); err != nil {
		return err
	}

	userLoginCredential := models.LoginCredential{
		Email:    apiUser.Email,
		Password: apiUser.Password,
	}

	err := repo.CreateHashedUserCredential(userLoginCredential)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(repo dao.Repo, apiUpdateUser models.ApiUpdateUser, email string) error {
	if err := repo.UpdateUser(apiUpdateUser, email); err != nil {
		return err
	}
	return nil
}

func GetAllUsers(repo dao.Repo) (apiUsers models.ApiUsers, err error) {

	if apiUsers, err = repo.GetAllUsers(); err != nil {
		return apiUsers, err
	}
	return apiUsers, nil
}

func validateUser(apiUser models.ApiCreateUser) string {

	var errs []string
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(apiUser)
	if err == nil {
		return ""
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr.Error())
	}
	return strings.Join(errs, ", ")
}
