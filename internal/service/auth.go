package service

import (
	"Start/internal/repository"
	"Start/internal/shared/utils"
	"Start/internal/store"
	"Start/internal/types"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) SignUp(input types.SignUpInput) (*store.User, error) {
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

	role := input.Role
	if role != "admin" {
		role = "user"
	}

	user := &store.User{
		ID:           uuid.NewString(),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPwd),
		Role:         role,
		Status:       "active",
		CreatedAt:    time.Now(),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(input types.LoginInput) (*types.LoginResponse, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	access, refresh, err := utils.GenerateTokens(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		User: types.UserDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		},
	}, nil
}

func (s *authService) RefreshToken(refreshToken string) (*types.TokenPair, error) {
	claims, err := utils.VerifyToken(refreshToken, utils.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	userID := claims["userId"].(string)
	email := claims["email"].(string)
	role := claims["role"].(string)

	accessToken, refreshToken, err := utils.GenerateTokens(userID, email, role)
	if err != nil {
		return nil, err
	}

	return &types.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) ChangePassword(userID, currentPassword, newPassword string) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(currentPassword)) != nil {
		return errors.New("incorrect current password")
	}

	hashedNew, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.UpdatePassword(userID, string(hashedNew))
}
