package middleware

import (
	"github.com/gin-gonic/gin"

	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		log.Printf("[GIN] %s %s | %d | %s",
			c.Request.Method,
			c.Request.URL.Path,
			status,
			latency)

	}
}
