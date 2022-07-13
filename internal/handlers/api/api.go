package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/indigowar/blog-site/internal/handlers/api/v1"
)

// Init api handlers
func Init(r *gin.RouterGroup) {
	v1.Init(r.Group("/v1"))
}
