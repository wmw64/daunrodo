package v1

import (
	"daunrodo/internal/entity"
	"daunrodo/internal/usecase"
	"daunrodo/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type downloadRoutes struct {
	u usecase.DownloadService
	l logger.ILogger
}

func newDLRoutes(handler *gin.RouterGroup, u usecase.DownloadService, l logger.ILogger) {
	r := &downloadRoutes{u, l}

	h := handler.Group("")
	{
		h.GET("/link", r.dl)
	}
}

type downloadResponse struct {
	File []entity.File `json:"file"`
}

func (r *downloadRoutes) dl(c *gin.Context) {
	link := c.Query("url")
	files, err := r.u.Download(c.Request.Context(), link)
	if err != nil {
		r.l.Error(err, "http - v1 - dl")
		c.AbortWithStatusJSON(http.StatusInternalServerError, "500 Internal Server Error")

		return
	}

	c.JSON(http.StatusOK, downloadResponse{files})
}
