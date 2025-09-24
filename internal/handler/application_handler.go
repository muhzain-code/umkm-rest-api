package handler

import (
	"net/http"
	"strconv"
	"umkm-api/internal/request"
	"umkm-api/internal/service"
	"umkm-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type ApplicationHandler struct {
	service service.ApplicationService
}

func NewApplicationHandler(service service.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{service: service}
}

func (y *ApplicationHandler) CreateApplication(ctx *gin.Context) {
	var req request.CreateApplicationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	app, err := y.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, app)
}

func (h *ApplicationHandler) GetAllApplication(ctx *gin.Context) {
	// ambil query params
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("per_page", "25")
	name := ctx.Query("name")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	// panggil service
	result, err := h.service.GetAll(page, limit, name)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	response.SuccessWithMeta(ctx, "Success fetch applications", &result.Meta, result.Data)
}

func (h *ApplicationHandler) GetApplicationByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	app, err := h.service.GetByID(id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	response.Success(ctx, "Success fetch application", app)
}

func (h *ApplicationHandler) UpdateApplication(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req request.UpdateApplicationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app, err := h.service.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h *ApplicationHandler) DeleteApplication(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}
