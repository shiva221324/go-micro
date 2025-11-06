package service

import (
	"auth-service/internal/model"
	"auth-service/internal/repository"
	"auth-service/pkg/jwt"
	"context"
)

type AuthService interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, email, password string) (string, *model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(ctx context.Context, user *model.User) (*model.User, error) {

	_, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return nil, ErrEmailExists
	}
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *authService) Login(ctx context.Context, email, password string) (string, *model.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}

	if err := user.CheckPassword(password); err != nil {
		return "", nil, ErrInvalidCredentials
	}
	token, err := jwt.GenerateToken(user.ID.String(), user.Email, user.Role, "24h")
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}

func (s *authService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
