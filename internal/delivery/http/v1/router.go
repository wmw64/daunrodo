package v1

import (
	"net/http"

	"daunrodo/internal/usecase"
	"daunrodo/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.ILogger, u usecase.DownloadService) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Routers
	h := handler.Group("")
	{
		newDLRoutes(h, u, l)
	}
}
