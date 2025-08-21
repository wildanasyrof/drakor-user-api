package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/dto"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"github.com/wildanasyrof/drakor-user-api/internal/repository"
	"github.com/wildanasyrof/drakor-user-api/pkg/hash"
	"github.com/wildanasyrof/drakor-user-api/pkg/jwt"
)

type AuthService interface {
	Register(req *dto.RegisterUserRequest) (*dto.UserResponse, *dto.TokenResponse, error)
	Login(req *dto.LoginUserRequest) (*dto.UserResponse, *dto.TokenResponse, error)
	Refresh(refreshToken string) (string, error)
	Logout(refreshToken string) error
	LogoutAll(uid uuid.UUID) error
}

type authService struct {
	userRepository  repository.UserRepository
	tokenRepository repository.TokenRepository
	jwtService      jwt.JWTService
}

func NewAuthService(userRepository repository.UserRepository, tokenRepository repository.TokenRepository, jwtService jwt.JWTService) AuthService {
	return &authService{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		jwtService:      jwtService,
	}
}

// Login implements AuthService.
func (a *authService) Login(req *dto.LoginUserRequest) (*dto.UserResponse, *dto.TokenResponse, error) {
	user, err := a.userRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, nil, err
	}

	if err := hash.ComparePassword(user.Password, req.Password); err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	token, err := a.GenerateToken(user.ID)
	if err != nil {
		return nil, nil, err
	}

	return &dto.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			ImgUrl:   user.ImgUrl,
		}, &dto.TokenResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}, nil
}

// Register implements AuthService.
func (a *authService) Register(req *dto.RegisterUserRequest) (*dto.UserResponse, *dto.TokenResponse, error) {
	hashedPassword := hash.HashPassword(req.Password)

	user := &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := a.userRepository.Create(user); err != nil {
		return nil, nil, err
	}

	token, err := a.GenerateToken(user.ID)
	if err != nil {
		return nil, nil, err
	}

	return &dto.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			ImgUrl:   user.ImgUrl,
		}, &dto.TokenResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}, nil
}

// Refresh implements AuthService.
func (a *authService) Refresh(refreshToken string) (string, error) {
	token, err := a.tokenRepository.Find(refreshToken)
	if err != nil || token.Revoked || time.Now().After(token.ExpiresAt) {
		return "", errors.New("invalid or expired refresh token")
	}

	newToken, err := a.jwtService.GenerateAccessToken(token.UserID)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

// Logout implements AuthService.
func (a *authService) Logout(refreshToken string) error {
	return a.tokenRepository.Revoke(refreshToken)
}

// LogoutAll implements AuthService.
func (a *authService) LogoutAll(uid uuid.UUID) error {
	return a.tokenRepository.RevokeAllForUser(uid)
}

func (a *authService) GenerateToken(userID uuid.UUID) (*dto.TokenResponse, error) {
	accessToken, err := a.jwtService.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.jwtService.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	refreshTokenEntity := &entity.RefreshToken{
		Token:     refreshToken,
		UserID:    userID,
		ExpiresAt: time.Now().Add(a.jwtService.GetRefreshTokenDuration()),
	}
	if err := a.tokenRepository.Create(refreshTokenEntity); err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
