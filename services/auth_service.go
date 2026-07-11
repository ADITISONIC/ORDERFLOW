package services

import (
	"errors"
	"orderflow/models"
	"orderflow/repositories"
	"orderflow/utils"
)

func Register(name, email, password string) error {

	// Check if user already exists
	existingUser, _ := repositories.GetUserByEmail(email)

	if existingUser.ID != 0 {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return repositories.CreateUser(&user)
}

func Login(email, password string) (string, error) {

	user, err := repositories.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}