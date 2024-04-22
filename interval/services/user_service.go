package services

import (
	"context"
	"fmt"
	"go-jwt/interval/models"
	"go-jwt/interval/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *userService) GetUserById(ctx context.Context, userID int) (*models.User, error) {
	user, err := s.userRepo.GetById(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user with ID %d not found", userID)
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *userService) DeleteUser(ctx context.Context, userID int) error {
	return s.userRepo.Delete(ctx, userID)
}
