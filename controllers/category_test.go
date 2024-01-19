package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/olooeez/video-vault/database"
	"gitlab.com/olooeez/video-vault/models"
)

func SetupTest() *gin.Engine {
	database.ConnectForTest()
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TeardownTest() {
	database.CloseForTest()
}

func CreateCategoryMock(category models.Category) int {
	database.DB.Create(&category)
	return int(category.ID)
}

func DeleteCategoryMock(category models.Category, id int) {
	database.DB.Delete(&category, id)
}

func TestGetCategories(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.GET("/api/v1/categories", GetCategories)

	category := models.Category{Title: "Category 1", Color: "#FFF"}
	id := CreateCategoryMock(category)
	defer DeleteCategoryMock(category, id)

	req, _ := http.NewRequest("GET", "/api/v1/categories", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var categories []models.Category
	err := json.Unmarshal(res.Body.Bytes(), &categories)

	assert.Nil(t, err, "error unmarshalling response body")
	assert.Equal(t, http.StatusOK, res.Code)

	var found bool
	for _, c := range categories {
		if c.Title == category.Title && c.Color == category.Color {
			found = true
			break
		}
	}

	assert.True(t, found, "category not found in the response")
}

func TestGetCategory(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.GET("/api/v1/categories/:id", GetCategory)

	category := models.Category{Title: "Test Category", Color: "#ABC"}
	id := CreateCategoryMock(category)
	defer DeleteCategoryMock(category, id)

	url := "/api/v1/categories/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var retrievedCategory models.Category
	err := json.Unmarshal(res.Body.Bytes(), &retrievedCategory)
	assert.Nil(t, err, "error unmarshalling response body")

	assert.Equal(t, category.Title, retrievedCategory.Title)
	assert.Equal(t, category.Color, retrievedCategory.Color)
}

func TestCreateCategory(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.POST("/api/v1/categories", CreateCategory)

	newCategory := models.Category{Title: "New Category", Color: "#123"}
	requestBody, err := json.Marshal(newCategory)
	assert.Nil(t, err, "error marshalling request body")

	req, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var createdCategory models.Category
	err = json.Unmarshal(res.Body.Bytes(), &createdCategory)
	assert.Nil(t, err, "error unmarshalling response body")

	assert.Equal(t, newCategory.Title, createdCategory.Title)
	assert.Equal(t, newCategory.Color, createdCategory.Color)

	defer DeleteCategoryMock(createdCategory, int(createdCategory.ID))
}

func TestUpdateCategory(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.PUT("/api/v1/categories/:id", UpdateCategory)

	originalCategory := models.Category{Title: "Original Category", Color: "#999"}
	id := CreateCategoryMock(originalCategory)
	defer DeleteCategoryMock(originalCategory, id)

	updatedCategory := models.Category{Title: "Updated Category", Color: "#456"}
	updatedCategory.ID = uint(id)

	requestBody, err := json.Marshal(updatedCategory)
	assert.Nil(t, err, "Error marshalling request body")

	url := "/api/v1/categories/" + strconv.Itoa(id)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	updatedCategoryFromDB := models.Category{}
	database.DB.First(&updatedCategoryFromDB, id)

	assert.Equal(t, updatedCategory.Title, updatedCategoryFromDB.Title)
	assert.Equal(t, updatedCategory.Color, updatedCategoryFromDB.Color)
}

func TestDeleteCategory(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.DELETE("/api/v1/categories/:id", DeleteCategory)

	categoryToDelete := models.Category{Title: "Category to Delete", Color: "#777"}
	id := CreateCategoryMock(categoryToDelete)

	url := "/api/v1/categories/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", url, nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	deletedCategoryFromDB := models.Category{}
	database.DB.First(&deletedCategoryFromDB, id)

	assert.Equal(t, uint(0), deletedCategoryFromDB.ID)
}

func TestGetCategoryVideos(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.GET("/api/v1/categories/:id/videos", GetCategoryVideos)

	category := models.Category{Title: "Category with Videos", Color: "#555"}
	categoryID := CreateCategoryMock(category)
	defer DeleteCategoryMock(category, categoryID)

	video1 := models.Video{Title: "Video 1", CategoryID: uint(categoryID)}
	video2 := models.Video{Title: "Video 2", CategoryID: uint(categoryID)}

	database.DB.Create(&video1)
	database.DB.Create(&video2)
	defer database.DB.Delete(&video1, video1.ID)
	defer database.DB.Delete(&video2, video2.ID)

	url := "/api/v1/categories/" + strconv.Itoa(categoryID) + "/videos"
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var videos []models.Video
	err := json.Unmarshal(res.Body.Bytes(), &videos)
	assert.Nil(t, err, "Error unmarshalling response body")

	for _, v := range videos {
		assert.Equal(t, uint(categoryID), v.CategoryID)
	}
}
