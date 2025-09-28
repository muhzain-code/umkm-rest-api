package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"umkm-api/internal/repository/filter"
	"umkm-api/internal/request"
	"umkm-api/internal/service"
	"umkm-api/pkg/response"
	"umkm-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Photo != nil {
		filename := uuid.New().String() + filepath.Ext(req.Photo.Filename)
		savePath := filepath.Join("uploads", "categories", filename)
		photo := "categories/" + filename

		uploadDir := filepath.Join("uploads", "categories")

		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			response.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("failed to create upload directory: %w", err))
			return
		}

		if err := ctx.SaveUploadedFile(req.Photo, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
			return
		}
		req.PhotoPath = &photo
	}

	category, err := h.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

func (h *CategoryHandler) GetAllCategory(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("per_page", "25")
	name := ctx.Query("name")
	statusStr := ctx.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	var status *bool
	if statusStr != "" {
		val := statusStr == "true"
		status = &val
	}

	filter := filter.CategoryFilter{
		Name:     name,
		IsActive: status,
	}

	result, err := h.service.GetAll(page, limit, filter)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
	}

	for i := range result.Data {
		result.Data[i].Photo = utils.URL(ctx, result.Data[i].Photo)
	}

	response.SuccessWithMeta(ctx, "Success fetch categories", &result.Meta, result.Data)

}

func (h *CategoryHandler) GetCategoryByID(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	category.Photo = utils.URL(ctx, category.Photo)

	response.Success(ctx, "Success fetch category", category)
}

func (h *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var req request.UpdateCategoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var filename *string
	if req.Photo != nil {
		newPhoto := uuid.New().String() + filepath.Ext(req.Photo.Filename)
		newPath := filepath.Join("uploads", "categories", newPhoto)
		photo := "categories/" + newPhoto

		if err := ctx.SaveUploadedFile(req.Photo, newPath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
			return
		}

		filename = &photo
	}

	category, err := h.service.Update(id, req, filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category deleted succesfully"})
}
