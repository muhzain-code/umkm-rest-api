package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Meta struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	LastPage    int `json:"last_page"`
	From        int `json:"from"`
	To          int `json:"to"`
}

func SuccessWithMeta(ctx *gin.Context, message string, meta *Meta, data interface{}) {
	response := gin.H{
		"message": message,
		"data":    data,
	}
	if meta != nil {
		response["meta"] = meta
	}

	ctx.JSON(http.StatusOK, response)
}

func Success(ctx *gin.Context, message string, data interface{}) {
	response := gin.H{
		"message": message,
		"data":    data,
	}

	ctx.JSON(http.StatusOK, response)

}

func ErrorResponse(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, gin.H{
		"message": err.Error(),
	})
}
