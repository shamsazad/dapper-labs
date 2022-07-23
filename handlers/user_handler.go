package handlers

import (
	"dapper-labs/auth"
	"dapper-labs/dao"
	"dapper-labs/models"
	"dapper-labs/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthenticationToken struct {
	Token string `json:"token"`
}

func SignUp(DAO dao.DaoInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var apiUser models.ApiCreateUser
		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&apiUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}

		err = service.UserSignUp(DAO, apiUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}

		tokenStr, err := auth.GenerateJWT(apiUser.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
		}

		authToken := AuthenticationToken{
			Token: tokenStr,
		}
		jsonResp, err := json.Marshal(authToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResp)
		return
	}
}

func UpdateUser(DAO dao.DaoInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var apiUpdateUser models.ApiUpdateUser

		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&apiUpdateUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}
		email := r.Context().Value("email")
		emailString := fmt.Sprintf("%v", email)

		if err = service.UpdateUser(DAO, apiUpdateUser, emailString); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func GetAllUsers(DAO dao.DaoInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var apiUsers models.ApiUsers
		w.Header().Set("Content-Type", "application/json")

		apiUsers, err := service.GetAllUsers(DAO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}

		jsonResp, err := json.Marshal(apiUsers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		return
	}
}
