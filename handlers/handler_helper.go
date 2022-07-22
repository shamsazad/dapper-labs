package handlers

import (
	"dapper-labs/models"
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

type DapperError struct {
	Error string `json:"error"`
}

func InitializeError(errorMessage string) []byte {
	dapperError := DapperError{
		Error: errorMessage,
	}
	jsonResp, err := json.Marshal(dapperError)
	if err != nil {
		return []byte(err.Error())
	}

	return jsonResp
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
