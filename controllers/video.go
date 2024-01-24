package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/video-vault/database"
	"gitlab.com/olooeez/video-vault/models"
)

func GetVideos(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		handleBadRequest(c, errors.New("invalid page number"))
		return
	}

	videosPerPage := 10
	offset := (page - 1) * videosPerPage

	var videos []models.Video
	database.DB.Offset(offset).Limit(videosPerPage).Find(&videos)

	c.JSON(http.StatusOK, videos)
}

func GetVideo(c *gin.Context) {
	video, err := findVideoByID(c)
	if err != nil {
		handleNotFoundError(c, "Video not found")
		return
	}

	c.JSON(http.StatusOK, video)
}

func CreateVideo(c *gin.Context) {
	var video models.Video

	if err := c.ShouldBindJSON(&video); err != nil {
		handleBadRequest(c, err)
		return
	}

	if err := models.ValidateVideoData(&video); err != nil {
		handleBadRequest(c, err)
		return
	}

	database.DB.Create(&video)
	c.JSON(http.StatusCreated, video)
}

func UpdateVideo(c *gin.Context) {
	video, err := findVideoByID(c)
	if err != nil {
		handleNotFoundError(c, "Video not found")
		return
	}

	if err := c.ShouldBindJSON(&video); err != nil {
		handleBadRequest(c, err)
		return
	}

	if err := models.ValidateVideoData(&video); err != nil {
		handleBadRequest(c, err)
		return
	}

	database.DB.Save(&video)
	c.JSON(http.StatusOK, video)
}

func DeleteVideo(c *gin.Context) {
	video, err := findVideoByID(c)
	if err != nil {
		handleNotFoundError(c, "Video not found")
		return
	}

	database.DB.Delete(&video)
	c.JSON(http.StatusOK, gin.H{
		"message": "Video deleted successfully",
		"status":  http.StatusOK,
	})
}

func SearchVideos(c *gin.Context) {
	var videos []models.Video
	query := c.Query("query")
	database.DB.Where("title LIKE ?", "%"+query+"%").Find(&videos)
	c.JSON(http.StatusOK, videos)
}

func findVideoByID(c *gin.Context) (models.Video, error) {
	var video models.Video
	id := c.Param("id")
	err := database.DB.First(&video, id).Error
	return video, err
}
