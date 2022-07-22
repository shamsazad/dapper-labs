package middleware

import (
	"context"
	"dapper-labs/auth"
	"dapper-labs/handlers"
	"net/http"
)

func AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-Type", "application/json")
			tokenString := r.Header.Get("x-authentication-token")
			if tokenString == "" {
				errorMsg := handlers.InitializeError("request does not contain an access token")
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write(errorMsg)
				return
			}
			email, err := auth.ValidateToken(tokenString)
			if err != nil {
				errorMsg := handlers.InitializeError(err.Error())
				rw.WriteHeader(http.StatusUnauthorized)
				rw.Write(errorMsg)
				return
			}
			ctx := context.WithValue(r.Context(), "email", email)
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
