package handler

import (
	"net/http"
	"umkm-api/internal/request"
	"umkm-api/internal/service"

	"github.com/gin-gonic/gin"
)

type LogHistoryHandler struct {
	service     service.LogHistoryService
	activityLog service.ActivityLogService
}

func NewLogHistoryHandler(service service.LogHistoryService, activityLog service.ActivityLogService) *LogHistoryHandler {
	return &LogHistoryHandler{service: service, activityLog: activityLog}
}

func (h *LogHistoryHandler) Create(c *gin.Context) {
	var req request.LogHistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.IPAddress = c.ClientIP()
	req.UserAgent = c.Request.UserAgent()

	_, err := h.service.ValidateProduct(req.ProductID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	h.service.CreateAsync(req)

	h.activityLog.Log(c, "buy", req.ProductID.String())

	c.JSON(http.StatusCreated, gin.H{
		"message": "Log is being recorded",
	})
}
