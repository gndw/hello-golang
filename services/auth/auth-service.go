package auth

import (
	"fmt"
	"hello-golang/errors"
	"hello-golang/models"
	"hello-golang/services/database"
)

type Auth struct {
	Token string `json:"token"`
}

func Login(username string, password string) (*Auth, *errors.Error) {

	if username != "username" || password != "password" {
		return nil, errors.CreateError(errors.WrongUserAndPassword)
	}

	signedToken, token_error := GenerateToken("123")
	if token_error != nil {
		return nil, token_error
	}

	return &Auth{Token: signedToken}, nil
}

func Register(email string, username string, password string) (*Auth, *errors.Error) {

	user := &models.User{
		Email:    email,
		Username: username,
		Password: password,
	}
	_, err := database.Db.Model(user).Insert()
	if err != nil {
		panic(err)
	}

	// Select all users.
	var users []models.User
	err = database.Db.Model(&users).Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	signedToken, token_error := GenerateToken("123")
	if token_error != nil {
		return nil, token_error
	}

	return &Auth{Token: signedToken}, nil
}
