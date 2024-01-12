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

func GetVideo(c *gin.Context) {
	var video models.Video
	id := c.Param("id")
	database.DB.First(&video, id)

	if video.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Video not found",
			"status":  http.StatusNotFound,
		})

		return
	}

	c.JSON(http.StatusOK, video)
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
	c.JSON(http.StatusCreated, video)
}

func UpdateVideo(c *gin.Context) {
	var video models.Video
	id := c.Param("id")
	database.DB.First(&video, id)

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

	database.DB.Model(&video).UpdateColumns(video)
	c.JSON(http.StatusOK, video)
}

func DeleteVideo(c *gin.Context) {
	var video models.Video
	id := c.Param("id")
	database.DB.First(&video, id)

	if video.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Video not found",
			"status":  http.StatusNotFound,
		})

		return
	}

	database.DB.Delete(&video)
	c.JSON(http.StatusOK, gin.H{
		"message": "Video deleted successfully",
		"status":  http.StatusOK,
	})
}
