package controllers_test

import (
	"bytes"
	"encoding/json"
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

func SetupTest() *gin.Engine {
	database.ConnectForTest()
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TeardownTest() {
	database.CloseForTest()
}

func CreateVideoMock(video models.Video) int {
	database.DB.Create(&video)
	return int(video.ID)
}

func DeleteVideoMock(video models.Video, id int) {
	database.DB.Delete(&video, id)
}

func TestGetVideos(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.GET("/api/v1/videos", controllers.GetVideos)

	video := models.Video{Title: "Video 1", Description: "Description 1", URL: "http://url1.com"}
	id := CreateVideoMock(video)
	defer DeleteVideoMock(video, id)

	req, _ := http.NewRequest("GET", "/api/v1/videos", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var videos []models.Video
	err := json.Unmarshal(res.Body.Bytes(), &videos)

	assert.Nil(t, err, "error unmarshalling response body")
	assert.Equal(t, http.StatusOK, res.Code)

	var found bool
	for _, v := range videos {
		if v.Title == video.Title && v.Description == video.Description && v.URL == video.URL {
			found = true
			break
		}
	}

	assert.True(t, found, "video not found in response")
}

func TestGetVideo(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.GET("/api/v1/videos/:id", controllers.GetVideo)

	video := models.Video{Title: "Video 2", Description: "Description 2", URL: "http://url2.com"}
	id := CreateVideoMock(video)
	defer DeleteVideoMock(video, id)

	url := "/api/v1/videos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", url, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var videoGet models.Video
	err := json.Unmarshal(res.Body.Bytes(), &videoGet)
	assert.Nil(t, err, "error unmarshalling response body")

	assert.Equal(t, video.Title, videoGet.Title)
	assert.Equal(t, video.Description, videoGet.Description)
	assert.Equal(t, video.URL, videoGet.URL)
}

func TestCreateVideo(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.POST("/api/v1/videos", controllers.CreateVideo)

	video := models.Video{Title: "Video 3", Description: "Description 3", URL: "http://url3.com", CategoryID: 1}
	jsonValue, err := json.Marshal(video)
	assert.Nil(t, err, "error marshalling request body")

	req, _ := http.NewRequest("POST", "/api/v1/videos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)

	var createdVideo models.Video
	err = json.Unmarshal(res.Body.Bytes(), &createdVideo)
	assert.Nil(t, err, "error unmarshalling response body")

	assert.Equal(t, video.Title, createdVideo.Title)
	assert.Equal(t, video.Description, createdVideo.Description)
	assert.Equal(t, video.URL, createdVideo.URL)

	defer DeleteVideoMock(createdVideo, int(createdVideo.ID))
}

func TestUpdateVideo(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.PUT("/api/v1/videos/:id", controllers.UpdateVideo)

	originalVideo := models.Video{Title: "Original Video", Description: "Original Description", URL: "http://original.com", CategoryID: 1}
	id := CreateVideoMock(originalVideo)
	defer DeleteVideoMock(originalVideo, id)

	updatedVideo := models.Video{Title: "Updated Video", Description: "Updated Description", URL: "http://updated.com", CategoryID: 1}
	updatedVideo.ID = uint(id)

	requestBody, err := json.Marshal(updatedVideo)
	assert.Nil(t, err, "error marshalling request body")

	url := "/api/v1/videos/" + strconv.Itoa(id)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	updatedVideoFromDB := models.Video{}
	database.DB.First(&updatedVideoFromDB, id)

	assert.Equal(t, updatedVideo.Title, updatedVideoFromDB.Title)
	assert.Equal(t, updatedVideo.Description, updatedVideoFromDB.Description)
	assert.Equal(t, updatedVideo.URL, updatedVideoFromDB.URL)
}

func TestDeleteVideo(t *testing.T) {
	t.Cleanup(TeardownTest)

	r := SetupTest()
	r.DELETE("/api/v1/videos/:id", controllers.DeleteVideo)

	videoToDelete := models.Video{Title: "Video to delete", Description: "Description to delete", URL: "http://delete.com"}
	id := CreateVideoMock(videoToDelete)

	url := "/api/v1/videos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", url, nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	deletedVideoFromDB := models.Video{}
	database.DB.First(&deletedVideoFromDB, id)

	assert.Equal(t, uint(0), deletedVideoFromDB.ID)
}
