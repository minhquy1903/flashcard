package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/minhquy1903/flashcard-api/internal/auth/model"
	"github.com/minhquy1903/flashcard-api/internal/auth/presenter"
	"github.com/minhquy1903/flashcard-api/internal/auth/repository"
	"github.com/minhquy1903/flashcard-api/pkg/token"
)

type AuthService struct {
	userRepo     repository.UserRepository
	tokenManager token.TokenManager
}

func NewAuthService(userRepo repository.UserRepository, tokenManager token.TokenManager) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (s *AuthService) Register(c context.Context, registerReq presenter.RegisterRequest) error {

	_, err := s.userRepo.FindByEmail(c, registerReq.Email)

	if err == nil {
		return errors.New("your email is already existed")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), 5)

	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &model.User{
		Email:    registerReq.Email,
		Name:     registerReq.Name,
		Password: string(hashedPassword),
	}

	return s.userRepo.Create(c, user)
}

func (s *AuthService) Login(c context.Context, loginReq presenter.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(c, loginReq.Email)
	if err != nil {
		return "", errors.New("your email or password is not valid")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return "", errors.New("your email or password is not valid")
	}

	token, err := s.tokenManager.NewJWT(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
