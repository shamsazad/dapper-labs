package handlers

import (
	"dapper-labs/auth"
	"dapper-labs/dao"
	"dapper-labs/models"
	"dapper-labs/service"
	"encoding/json"
	"net/http"
)

func Login(repo dao.DaoInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginCredential models.LoginCredential

		w.Header().Set("Content-Type", "application/json")
		err := json.NewDecoder(r.Body).Decode(&loginCredential)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}

		if err = service.Login(repo, loginCredential); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			errorMsg := InitializeError(err.Error())
			w.Write(errorMsg)
			return
		}

		tokenStr, err := auth.GenerateJWT(loginCredential.Email)
		authToken := AuthenticationToken{
			Token: tokenStr,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(authToken)
	}
}
