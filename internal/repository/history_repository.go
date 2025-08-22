package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wildanasyrof/drakor-user-api/internal/domain/entity"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	Create(req *entity.History) error
	DeleteByID(id uuid.UUID) error
	GetByID(userID uuid.UUID, historyID uuid.UUID) (*entity.History, error)
	GetByUserID(userID uuid.UUID) ([]entity.History, error)
	Update(req *entity.History) error
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{
		db: db,
	}
}

// Create implements HistoryRepository.
func (h *historyRepository) Create(req *entity.History) error {
	return h.db.Create(req).Error
}

// DeleteByID implements HistoryRepository.
func (h *historyRepository) DeleteByID(id uuid.UUID) error {
	return h.db.Delete(&entity.History{}, id).Error
}

// GetByID implements HistoryRepository.
func (h *historyRepository) GetByID(userID uuid.UUID, historyID uuid.UUID) (*entity.History, error) {
	var history entity.History
	if err := h.db.Where("user_id = ? AND id = ?", userID, historyID).First(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

// GetByUserID implements HistoryRepository.
func (h *historyRepository) GetByUserID(userID uuid.UUID) ([]entity.History, error) {
	var histories []entity.History
	if err := h.db.Where("user_id = ?", userID).Find(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

// Update implements HistoryRepository.
func (h *historyRepository) Update(req *entity.History) error {
	res := h.db.Model(&entity.History{}).Where("id = ?", req.ID).Updates(req)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no rows affected, history not found")
	}
	return nil
}
