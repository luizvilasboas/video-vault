package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/video-vault/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/v1/videos", controllers.GetVideos)
	r.GET("/api/v1/videos/:id", controllers.GetVideo)
	r.POST("/api/v1/videos", controllers.CreateVideo)
	r.Run()
}
