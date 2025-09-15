package handler

import (
	"github.com/google/uuid"
	"umkm-api/internal/request"
	"umkm-api/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UmkmHandler struct {
	service service.UmkmService
}

func NewUmkmHandler(service service.UmkmService) *UmkmHandler {
	return &UmkmHandler{service: service}
}

func (c *UmkmHandler) CreateUmkm(ctx *gin.Context) {
	var req request.CreateUmkmRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	umkm, err := c.service.Create(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusCreated, umkm)
}

func (c *UmkmHandler) GetAllUmkm(ctx *gin.Context) {
	umkms, err := c.service.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, umkms)
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

	ctx.JSON(http.StatusOK, umkm)
}

func (c *UmkmHandler) UpdateUmkm(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
	}

	var req request.UpdateUmkmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	umkm, err := c.service.Update(id, req)

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
