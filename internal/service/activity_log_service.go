// internal/service/activity_log_service.go
package service

import (
	"github.com/gin-gonic/gin"
	"umkm-api/internal/model"
	"umkm-api/internal/repository"
)

type ActivityLogService interface {
	Log(c *gin.Context, logType string, productID string)
	GetBySession(sessionID string) ([]model.ActivityLog, error)
}

type activityLogService struct {
	repo repository.ActivityLogRepository
}

func NewActivityLogService(repo repository.ActivityLogRepository) ActivityLogService {
	return &activityLogService{repo: repo}
}

func (s *activityLogService) Log(c *gin.Context, logType string, productID string) {
	ip := c.ClientIP()
	userAgent := c.Request.UserAgent()
	referrer := c.Request.Referer()

	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		sessionID = "guest"
	}

	log := model.ActivityLog{
		SessionID: &sessionID,
		ProductID: productID,
		LogType:   logType,
		IPAddress: &ip,
		UserAgent: &userAgent,
		Referrer:  &referrer,
	}

	go func() {
		_ = s.repo.Create(&log)
	}()
}

// Ambil log berdasarkan session
func (s *activityLogService) GetBySession(sessionID string) ([]model.ActivityLog, error) {
	return s.repo.FindBySessionID(sessionID)
}
