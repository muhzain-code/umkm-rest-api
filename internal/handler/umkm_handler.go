package handler

import (
	"umkm-api/internal/request"
	"umkm-api/internal/service"

	"github.com/google/uuid"

	"net/http"

	"strconv"
	"umkm-api/pkg/response"

	"github.com/gin-gonic/gin"
	"umkm-api/pkg/utils"
	"path/filepath"
)

type UmkmHandler struct {
	service service.UmkmService
}

func NewUmkmHandler(service service.UmkmService) *UmkmHandler {
	return &UmkmHandler{service: service}
}

func (h *UmkmHandler) CreateUmkm(ctx *gin.Context) {
    var req request.CreateUmkmRequest
    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // simpan file foto (jika ada)
    if req.PhotoProfile != nil {
        filename := uuid.New().String() + filepath.Ext(req.PhotoProfile.Filename)
        savePath := filepath.Join("uploads", "umkms", filename)
		photoProfile := "umkms/" + filename

        if err := ctx.SaveUploadedFile(req.PhotoProfile, savePath); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
            return
        }
        req.PhotoProfilePath = &photoProfile
    }

    umkm, err := h.service.Create(req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, umkm)
}


func (c *UmkmHandler) GetAllUmkm(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("per_page", "5")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	result, err := c.service.GetAll(page, limit)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	for i := range result.Data {
		result.Data[i].PhotoProfile = utils.URL(ctx, result.Data[i].PhotoProfile)
	}

	response.SuccessWithMeta(ctx, "Success fetch UMKMs", &result.Meta, result.Data)
}

func (c *UmkmHandler) GetUmkmByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	umkm, err := c.service.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	umkm.PhotoProfile = utils.URL(ctx, umkm.PhotoProfile)
	
	response.Success(ctx, "Success fetch umkm", umkm)
}


func (h *UmkmHandler) UpdateUmkm(ctx *gin.Context) {
    id, err := uuid.Parse(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    var req request.UpdateUmkmRequest
    if err := ctx.ShouldBind(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := req.Validate(); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Upload foto
    var fileName *string
    if req.PhotoProfile != nil {
        newFileName := uuid.NewString() + filepath.Ext(req.PhotoProfile.Filename)
        newPath := filepath.Join("uploads", "umkms", newFileName)
		photoProfile := "umkms/" + newFileName

        if err := ctx.SaveUploadedFile(req.PhotoProfile, newPath); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save photo"})
            return
        }

        fileName = &photoProfile
    }

    umkm, err := h.service.Update(id, req, fileName)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, umkm)
}


func (c *UmkmHandler) DeleteUmkm(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
	}

	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "umkm deleted succesfully"})
}
