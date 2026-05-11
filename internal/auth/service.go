package auth

import (
	"mangahub/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		ID:           username, // đơn giản hóa
		Username:     username,
		PasswordHash: string(hash),
	}

	return CreateUser(user)
}

func LoginUser(username, password string) (models.User, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return user, err
}