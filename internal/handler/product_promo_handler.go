package handler

import (
	"net/http"
	"strconv"
	"umkm-api/internal/repository/filter"
	"umkm-api/internal/request"
	"umkm-api/internal/service"
	"umkm-api/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductPromoHandler struct {
	service service.ProductPromoService
}

func NewProductPromoHandler(service service.ProductPromoService) *ProductPromoHandler {
	return &ProductPromoHandler{service: service}
}

func (h *ProductPromoHandler) CreateProductPromo(ctx *gin.Context) {
	var req request.CreateProductPromoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productPromo, err := h.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, productPromo)
}

func (h *ProductPromoHandler) GetAllProductPromo(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("per_page", "25")
	eventStr := ctx.Query("event_id")
	productStr := ctx.Query("product_id")
	statusStr := ctx.Query("status")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	eventID, _ := uuid.Parse(eventStr)
	productID, _ := uuid.Parse(productStr)

	var status *bool
	if statusStr != "" {
		val := statusStr == "true"
		status = &val
	}

	filter := filter.ProductPromoFilter{
		EventID:   eventID,
		ProductID: productID,
		IsActive:  status,
	}

	result, err := h.service.GetAll(page, limit, filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response.SuccessWithMeta(ctx, "success fetch product promos", &result.Meta, result.Data)
}

func (h *ProductPromoHandler) GetProductPromoByID(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	productPromo, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response.Success(ctx, "success fetch product promo", productPromo)
}

func (h *ProductPromoHandler) UpdateProductPromo(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	var req request.UpdateProductPromoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productPromo, err := h.service.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, productPromo)
}

func (h *ProductPromoHandler) DeleteProductPromo(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err := h.service.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product promo deleted successfully"})
}
