package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/video-vault/database"
	"gitlab.com/olooeez/video-vault/models"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Preload("Videos").Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.Preload("Videos").First(&category, id).Error; err != nil {
		handleNotFoundError(c, "Category not found")
		return
	}

	c.JSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		handleBadRequest(c, err)
		return
	}

	database.DB.Create(&category)
	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		handleNotFoundError(c, "Category not found")
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		handleBadRequest(c, err)
		return
	}

	if err := models.ValidateCategoryData(&category); err != nil {
		handleBadRequest(c, err)
		return
	}

	database.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.First(&category, id).Error; err != nil {
		handleNotFoundError(c, "Category not found")
		return
	}

	database.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
		"status":  http.StatusOK,
	})
}

func GetCategoryVideos(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := database.DB.Preload("Videos").First(&category, id).Error; err != nil {
		handleNotFoundError(c, "Category not found")
		return
	}

	c.JSON(http.StatusOK, category.Videos)
}

func handleBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
		"status":  http.StatusBadRequest,
	})
}

func handleNotFoundError(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": message,
		"status":  http.StatusNotFound,
	})
}
