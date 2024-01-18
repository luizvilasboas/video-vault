package controllers_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/olooeez/video-vault/controllers"
	"gitlab.com/olooeez/video-vault/database"
	"gitlab.com/olooeez/video-vault/models"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func CreateVideoMock(video models.Video) {
	database.DB.Create(&video)
	ID = int(video.ID)
}

func DeleteVideoMock(video models.Video) {
	database.DB.Delete(&video, ID)
}

func TestGetVideos(t *testing.T) {
	database.Connect()

	video := models.Video{Title: "Video 1", Description: "Description 1", URL: "http://url1.com"}
	CreateVideoMock(video)
	defer DeleteVideoMock(video)

	r := SetupTestRoutes()
	r.GET("/api/v1/videos", controllers.GetVideos)

	req, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var videos []models.Video
	json.Unmarshal(res.Body.Bytes(), &videos)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Video 1", videos[0].Title)
	assert.Equal(t, "Description 1", videos[0].Description)
	assert.Equal(t, "http://url1.com", videos[0].URL)
}

func TestGetVideo(t *testing.T) {
	database.Connect()

	video := models.Video{Title: "Video 2", Description: "Description 2", URL: "http://url2.com"}
	CreateVideoMock(video)
	defer DeleteVideoMock(video)

	r := SetupTestRoutes()
	r.GET("/api/v1/videos/:id", controllers.GetVideo)

	req, _ := http.NewRequest("GET", "/api/v1/videos/"+strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var videoGet models.Video
	json.Unmarshal(res.Body.Bytes(), &videoGet)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Video 2", videoGet.Title)
	assert.Equal(t, "Description 2", videoGet.Description)
	assert.Equal(t, "http://url2.com", videoGet.URL)
}

func TestCreateVideo(t *testing.T) {
	database.Connect()

	video := models.Video{Title: "Video 3", Description: "Description 3", URL: "http://url3.com", CategoryID: 1}

	r := SetupTestRoutes()
	r.POST("/api/v1/videos", controllers.CreateVideo)

	jsonValue, _ := json.Marshal(video)

	req, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewBuffer(jsonValue))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var videoGet models.Video
	json.Unmarshal(res.Body.Bytes(), &videoGet)

	log.Printf("%v\n", videoGet)

	ID = int(videoGet.ID)
	defer DeleteVideoMock(videoGet)

	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, "Video 3", videoGet.Title)
	assert.Equal(t, "Description 3", videoGet.Description)
	assert.Equal(t, "http://url3.com", videoGet.URL)
}

func TestUpdateVideo(t *testing.T) {
	database.Connect()

	video := models.Video{Title: "Video 4", Description: "Description 4", URL: "http://url4.com", CategoryID: 1}
	CreateVideoMock(video)
	defer DeleteVideoMock(video)

	r := SetupTestRoutes()
	r.PUT("/api/v1/videos/:id", controllers.UpdateVideo)

	video.ID = uint(ID)
	video.Title = "Video 5"
	video.Description = "Description 5"
	video.URL = "http://url5.com"

	jsonValue, _ := json.Marshal(video)

	req, _ := http.NewRequest("PUT", "/api/v1/videos/"+strconv.Itoa(ID), bytes.NewBuffer(jsonValue))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var videoGet models.Video
	json.Unmarshal(res.Body.Bytes(), &videoGet)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Video 5", videoGet.Title)
	assert.Equal(t, "Description 5", videoGet.Description)
	assert.Equal(t, "http://url5.com", videoGet.URL)
}

func TestDeleteVideo(t *testing.T) {
	database.Connect()

	video := models.Video{Title: "Video 6", Description: "Description 6", URL: "http://url6.com"}
	CreateVideoMock(video)

	r := SetupTestRoutes()
	r.DELETE("/api/v1/videos/:id", controllers.DeleteVideo)

	req, _ := http.NewRequest("DELETE", "/api/v1/videos/"+strconv.Itoa(ID), nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Contains(t, res.Body.String(), "Video deleted successfully")
}
