package repository

import (
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByUsername(username string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetByID(id string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id string) error {
	return r.db.Delete(&entity.User{}, id).Error
}
