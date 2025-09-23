package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"strconv"
	"umkm-api/internal/product/repository"
	"umkm-api/internal/product/request"
	"umkm-api/internal/product/service"
	"umkm-api/pkg/response"
	"umkm-api/pkg/utils"
)

const uploadDir = "uploads/products"

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req request.CreateProductRequest
	req.UmkmID = ctx.PostForm("umkm_id")

	categoryStr := ctx.PostForm("category_id")
	if categoryStr != "" {
		if cid, err := strconv.ParseInt(categoryStr, 10, 64); err == nil {
			req.CategoryID = cid
		} else {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid category_id"))
			return
		}
	}

	req.Name = ctx.PostForm("name")

	desc := ctx.PostForm("description")
	if desc != "" {
		req.Description = &desc
	}

	priceStr := ctx.PostForm("price")
	if priceStr != "" {
		if p, err := strconv.ParseFloat(priceStr, 64); err == nil {
			req.Price = p
		} else {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid price"))
			return
		}
	}

	req.Status = ctx.PostForm("status")

	photosJSON := ctx.PostForm("photos")
	if photosJSON != "" {
		var photos []request.PhotoRequest
		if err := json.Unmarshal([]byte(photosJSON), &photos); err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid photos json: %w", err))
			return
		}
		req.Photos = photos
	}

	marketplacesJSON := ctx.PostForm("marketplaces")
	if marketplacesJSON != "" {
		var mps []request.MarketplaceRequest
		if err := json.Unmarshal([]byte(marketplacesJSON), &mps); err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid marketplaces json: %w", err))
			return
		}
		req.Marketplaces = mps
	}

	if ctx.Request.MultipartForm != nil && ctx.Request.MultipartForm.File != nil {
		files := ctx.Request.MultipartForm.File["files"] // expect files[] multiple
		if len(files) > 0 {
			if len(req.Photos) == 0 {
				req.Photos = make([]request.PhotoRequest, len(files))
			}

			for i, fheader := range files {
				filename := uuid.New().String() + filepath.Ext(fheader.Filename)
				savePath := filepath.Join(uploadDir, filename)

				if err := ctx.SaveUploadedFile(fheader, savePath); err != nil {
					response.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("failed to save uploaded file: %w", err))
					return
				}

				if i < len(req.Photos) {
					req.Photos[i].FilePath = filepath.ToSlash(filepath.Join("products", filename))
				} else {
					req.Photos = append(req.Photos, request.PhotoRequest{
						FilePath: filepath.ToSlash(filepath.Join("products", filename)),
					})
				}
			}
		}
	}

	if req.UmkmID == "" {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("umkm_id is required"))
		return
	}
	if req.CategoryID == 0 {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("category_id is required"))
		return
	}
	if req.Name == "" {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("name is required"))
		return
	}
	if req.Price <= 0 {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("price is required and must be > 0"))
		return
	}
	if req.Status == "" {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("status is required"))
		return
	}

	product, err := h.service.Create(req)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("failed to create product: %w", err))
		return
	}

	response.Success(ctx, "Product created successfully", product)
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("per_page", "25")
	name := ctx.Query("name")
	statusStr := ctx.Query("status")
	umkmID := ctx.Query("umkm_id")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	var umkmUUID uuid.UUID
	if umkmID != "" {
		parsed, err := uuid.Parse(umkmID)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid umkm_id"))
			return
		}
		umkmUUID = parsed
	}

	var isActive *bool
	if statusStr != "" {
		val := statusStr == "true" || statusStr == "1"
		isActive = &val
	}

	filter := repository.ProductFilter{
		Name:     name,
		IsActive: isActive,
		UmkmID:   umkmUUID,
	}

	result, err := h.service.GetAll(page, limit, filter)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	for i := range result.Data {
		for j := range result.Data[i].Photos {
			fp := result.Data[i].Photos[j].FilePath
			url := utils.URL(ctx, &fp)
			if url != nil {
				result.Data[i].Photos[j].FilePath = *url 
			}
		}

		if result.Data[i].Category.Photo != nil {
			result.Data[i].Category.Photo = utils.URL(ctx, result.Data[i].Category.Photo)
		}
	}

	response.SuccessWithMeta(ctx, "Success fetch products", &result.Meta, result.Data)
}

func (h *ProductHandler) GetProductByID(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid product id"))
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}

	for i := range product.Photos {
		fp := product.Photos[i].FilePath
		url := utils.URL(ctx, &fp)
		if url != nil {
			product.Photos[i].FilePath = *url
		}
	}

	if product.Category.Photo != nil {
		product.Category.Photo = utils.URL(ctx, product.Category.Photo)
	}

	if product.Category.Photo != nil {
		product.Category.Photo = utils.URL(ctx, product.Category.Photo)
	}

	response.Success(ctx, "Success fetch product", product)
}

func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid product id"))
		return
	}

	var req request.UpdateProductRequest
	umkmIDStr := ctx.PostForm("umkm_id")
	if umkmIDStr != "" {
		parsedID, err := uuid.Parse(umkmIDStr)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid umkm_id"))
			return
		}
		req.UmkmID = parsedID
	}

	categoryStr := ctx.PostForm("category_id")
	if categoryStr != "" {
		if cid, err := strconv.ParseInt(categoryStr, 10, 64); err == nil {
			req.CategoryID = cid
		} else {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid category_id"))
			return
		}
	}

	req.Name = ctx.PostForm("name")
	desc := ctx.PostForm("description")
	if desc != "" {
		req.Description = &desc
	}

	priceStr := ctx.PostForm("price")
	if priceStr != "" {
		if p, err := strconv.ParseFloat(priceStr, 64); err == nil {
			req.Price = p
		} else {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid price"))
			return
		}
	}

	req.Status = ctx.PostForm("status")

	photosJSON := ctx.PostForm("photos")
	if photosJSON != "" {
		var photos []request.PhotoRequest
		if err := json.Unmarshal([]byte(photosJSON), &photos); err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid photos json: %w", err))
			return
		}
		req.Photos = photos
	}

	marketplacesJSON := ctx.PostForm("marketplaces")
	if marketplacesJSON != "" {
		var mps []request.MarketplaceRequest
		if err := json.Unmarshal([]byte(marketplacesJSON), &mps); err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid marketplaces json: %w", err))
			return
		}
		req.Marketplaces = mps
	}

	if ctx.Request.MultipartForm != nil && ctx.Request.MultipartForm.File != nil {
		files := ctx.Request.MultipartForm.File["files"]
		if len(files) > 0 {
			if len(req.Photos) == 0 {
				req.Photos = make([]request.PhotoRequest, len(files))
			}

			for i, fheader := range files {
				filename := uuid.New().String() + filepath.Ext(fheader.Filename)
				savePath := filepath.Join(uploadDir, filename)

				if err := ctx.SaveUploadedFile(fheader, savePath); err != nil {
					response.ErrorResponse(ctx, http.StatusInternalServerError, fmt.Errorf("failed to save uploaded file: %w", err))
					return
				}

				if i < len(req.Photos) {
					req.Photos[i].FilePath = filepath.ToSlash(filepath.Join("products", filename))
				} else {
					req.Photos = append(req.Photos, request.PhotoRequest{
						FilePath: filepath.ToSlash(filepath.Join("products", filename)),
					})
				}
			}
		}
	}

	product, err := h.service.Update(id, req)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	for i := range product.Photos {
		fp := product.Photos[i].FilePath
		url := utils.URL(ctx, &fp)
		if url != nil {
			product.Photos[i].FilePath = *url
		}
	}

	response.Success(ctx, "Product updated successfully", product)
}

func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid product id"))
		return
	}

	if err := h.service.Delete(id); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.Success(ctx, "Product deleted successfully", gin.H{"id": id})
}
