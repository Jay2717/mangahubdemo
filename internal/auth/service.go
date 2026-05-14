package auth

import (
	"mangahub/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Register
func (s *Service) RegisterUser(username, email, password string) error {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
	}

	return s.repo.CreateUser(user)
}

// login
func (s *Service) LoginUser(username, password string) (models.User, error) {

	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return user, err
	}

	return user, nil
}
