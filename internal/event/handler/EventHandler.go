package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"umkm-api/internal/event/request"
	"umkm-api/internal/event/service"
	"umkm-api/pkg/response"
	"umkm-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventHandler struct {
	service service.EventService
}

func NewEventHandler(service service.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (E *EventHandler) GetAllEvent(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("per_page", "25")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	result, err := E.service.GetAll(page, limit)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	for i := range result.Data {
		result.Data[i].Photo = utils.URL(ctx, result.Data[i].Photo)
	}

	response.SuccessWithMeta(ctx, "Success fetch Events", &result.Meta, result.Data)
}

func (D *EventHandler) GetEventByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := D.service.GetByID(int(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	event.Photo = utils.URL(ctx, event.Photo)
	response.Success(ctx, "Success fetch event", event)
}

func (h *EventHandler) CreateEvent(ctx *gin.Context) {
	var req request.CreateEventRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Photo != nil {
		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(req.Photo.Filename))
		savePath := filepath.Join("uploads", "events", filename) // folder khusus event
		photoPath := "events/" + filename                        // path relatif yang disimpan di DB

		// simpan file ke folder uploads/events
		if err := ctx.SaveUploadedFile(req.Photo, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
			return
		}
		req.PhotoPath = &photoPath
	}

	// simpan event ke DB lewat service
	event, err := h.service.CreateEvent(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

func (h *EventHandler) UpdateEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req request.UpdateEventRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// handle upload file (opsional)
	if req.Photo != nil {
		newFileName := uuid.NewString() + filepath.Ext(req.Photo.Filename)
		newPath := filepath.Join("uploads", "events", newFileName)
		photoPath := "events/" + newFileName

		if err := ctx.SaveUploadedFile(req.Photo, newPath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
			return
		}

		req.PhotoPath = &photoPath
	}

	event, err := h.service.UpdateEvent(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) DeleteEvent(ctx *gin.Context) {
    idParam := ctx.Param("id")

    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := h.service.DeleteEvent(id); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}

