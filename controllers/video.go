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

func CreateVideo(c *gin.Context) {
	var video models.Video

	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})

		return
	}

	if err := models.ValidateVideoData(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})

		return
	}

	database.DB.Create(&video)
	c.JSON(http.StatusOK, video)
}
