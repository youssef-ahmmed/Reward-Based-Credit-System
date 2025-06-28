package user

import (
	"Start/internal/shared/utils"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Service interface {
	SignUp(input SignUpInput) (*User, error)
	Login(input LoginInput) (*LoginResponse, error)
	RefreshToken(refreshToken string) (*TokenPair, error)
	ChangePassword(userID, currentPassword, newPassword string) error
	GetProfile(userID string) (*UserDTO, error)
	UpdateProfile(userID string, input UpdateProfileRequest) error
	GetWallet(userID string) (*Wallet, error)
	AddToWallet(userID string, credits int, points int) error
	DeductPointsTx(tx *gorm.DB, userID string, points int) error
}

type service struct {
	repo *Repository
}

func NewService(repo *Repository) Service {
	return &service{repo: repo}
}

func (s *service) SignUp(input SignUpInput) (*User, error) {
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

	user := &User{
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

func (s *service) Login(input LoginInput) (*LoginResponse, error) {
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

	return &LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		User: UserDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		},
	}, nil
}

func (s *service) RefreshToken(refreshToken string) (*TokenPair, error) {
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

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) ChangePassword(userID, currentPassword, newPassword string) error {
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

func (s *service) GetProfile(userID string) (*UserDTO, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (s *service) UpdateProfile(userID string, input UpdateProfileRequest) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	if input.Username != nil && *input.Username != user.Username {
		taken, err := s.repo.IsUsernameTaken(*input.Username)
		if err != nil {
			return err
		}
		if taken {
			return errors.New("username already exists")
		}
		user.Username = *input.Username
	}

	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		user.LastName = *input.LastName
	}

	return s.repo.UpdateUser(user)
}

func (s *service) GetWallet(userID string) (*Wallet, error) {
	wallet, err := s.repo.GetWalletByUserID(userID)
	if err == nil {
		return wallet, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newWallet := &Wallet{
			ID:             uuid.NewString(),
			UserID:         userID,
			CreditsBalance: 0,
			PointsBalance:  0,
			UpdatedAt:      time.Now(),
		}
		if err := s.repo.CreateWallet(newWallet); err != nil {
			return nil, err
		}
		return newWallet, nil
	}

	return nil, err
}

func (s *service) AddToWallet(userID string, credits int, points int) error {
	wallet, err := s.repo.GetWalletByUserID(userID)
	if err != nil {
		return err
	}

	wallet.CreditsBalance += credits
	wallet.PointsBalance += points
	wallet.UpdatedAt = time.Now()

	return s.repo.UpdateWallet(wallet)
}

func (s *service) DeductPointsTx(tx *gorm.DB, userID string, points int) error {
	return s.repo.DeductPointsTx(tx, userID, points)
}
