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
	database.DB.Preload("Videos").First(&category, id)

	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
			"status":  http.StatusNotFound,
		})

		return
	}

	c.JSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})

		return
	}

	database.DB.Create(&category)
	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")
	database.DB.First(&category, id)

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})

		return
	}

	if err := models.ValidateCategoryData(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})

		return
	}

	database.DB.Model(&category).UpdateColumns(category)
	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")
	database.DB.First(&category, id)

	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
			"status":  http.StatusNotFound,
		})

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
	database.DB.Preload("Videos").First(&category, id)
	c.JSON(http.StatusOK, category.Videos)
}
