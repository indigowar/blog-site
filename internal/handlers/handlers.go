package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/indigowar/blog-site/internal/handlers/api"
	"net/http"
)

// Init handlers
func Init(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, world",
		})
	})

	api.Init(r.Group("/api"))
}
