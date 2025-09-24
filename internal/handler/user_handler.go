package handler

import (
	"net/http"
	"umkm-api/internal/auth"
	"umkm-api/internal/request"
	"umkm-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service    service.UserService
	jwtService auth.JWTService
}

func NewHandler(service service.UserService, jwtService auth.JWTService) *UserHandler {
	return &UserHandler{service, jwtService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req *request.UserRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, _ := h.jwtService.GenerateToken(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
