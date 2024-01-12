package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/video-vault/database"
	"gitlab.com/olooeez/video-vault/models"
)

func GetVideos(c *gin.Context) {
	var videos []models.Video
	database.DB.Find(&videos)
	c.JSON(http.StatusOK, videos)
}
