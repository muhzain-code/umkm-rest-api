package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func URL(ctx *gin.Context, photo *string) *string {
    if photo == nil {
        return nil
    }

    scheme := "http"
    if ctx.Request.TLS != nil {
        scheme = "https"
    }

    baseURL := fmt.Sprintf("%s://%s/uploads/%s", scheme, ctx.Request.Host, *photo)
    return &baseURL
}