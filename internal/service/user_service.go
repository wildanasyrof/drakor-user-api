package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/dto"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"github.com/wildanasyrof/drakor-user-api/internal/repository"
	"github.com/wildanasyrof/drakor-user-api/pkg/hash"
)

type UserService interface {
	Get(userID uuid.UUID) (*entity.User, error)
	Update(userID uuid.UUID, req *dto.UpdateUserRequest) (*entity.User, error)
	UpdateAvatar(userID uuid.UUID, url string) (*entity.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Get implements UserService.
func (u *userService) Get(userID uuid.UUID) (*entity.User, error) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update implements UserService.
func (u *userService) Update(userID uuid.UUID, req *dto.UpdateUserRequest) (*entity.User, error) {
	user, err := u.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if req.Email != nil {
		user.Email = *req.Email
	}

	if req.Username != nil {
		user.Username = *req.Username
	}

	if req.Password != nil {
		hashedPassword := hash.HashPassword(*req.Password)
		user.Password = hashedPassword
	}

	if req.ImgUrl != nil {
		user.ImgUrl = *req.ImgUrl
	}

	if err := u.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateAvatar(userID uuid.UUID, url string) (*entity.User, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	user.ImgUrl = url
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
