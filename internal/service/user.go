package service

import (
	"Start/internal/repository"
	"Start/internal/types"
	"errors"
)

type userService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetProfile(userID string) (*types.UserDTO, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &types.UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (s *userService) UpdateProfile(userID string, input types.UpdateProfileRequest) error {
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
