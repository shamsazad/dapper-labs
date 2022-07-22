package models

type LoginCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
