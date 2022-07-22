package models

type ApiCreateUser struct {
	Email     string `validate:"required,email"`
	FirstName string `validate:"required" json:"first_name"`
	LastName  string `validate:"required" json:"last_name"`
	Password  string `validate:"required" json:"password,omitempty"`
}

type ApiUsers []ApiCreateUser

func (apiUser ApiCreateUser) ConvertApiUserToUser() User {
	return User{
		Email:     apiUser.Email,
		FirstName: apiUser.FirstName,
		LastName:  apiUser.LastName,
	}
}

func (apiUsers ApiUsers) ConvertApiUsersToUsers() Users {
	var users []User
	for _, user := range apiUsers {
		user := user.ConvertApiUserToUser()
		users = append(users, user)
	}
	return users
}
