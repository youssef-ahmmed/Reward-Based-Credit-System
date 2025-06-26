package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type SignUpInput struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func (s *Service) SignUp(input SignUpInput) (*User, error) {
	taken, err := s.repo.IsEmailOrUsernameTaken(input.Email, input.Username)
	if err != nil {
		return nil, err
	}

	if taken {
		return nil, errors.New("email or username already exists")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:           uuid.NewString(),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPwd),
		Role:         "user",
		Status:       "active",
		CreatedAt:    time.Now(),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
