package repository

import (
	"umkm-api/internal/model"

	"gorm.io/gorm"
)

type ActivityLogRepository interface {
	Create(log *model.ActivityLog) error
	FindBySessionID(sessionID string) ([]model.ActivityLog, error)
}

type activityLogRepository struct {
	db *gorm.DB
}

func NewActivityLogRepository(db *gorm.DB) ActivityLogRepository {
	return &activityLogRepository{db: db}
}

func (r *activityLogRepository) Create(log *model.ActivityLog) error {
	return r.db.Create(log).Error
}

func (r *activityLogRepository) FindBySessionID(sessionID string) ([]model.ActivityLog, error) {
	var logs []model.ActivityLog
	err := r.db.Where("session_id = ?", sessionID).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}
