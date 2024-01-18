package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/video-vault/controllers"
)

func HandleRequests() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/videos", controllers.GetVideos)
		v1.GET("/videos/:id", controllers.GetVideo)
		v1.POST("/videos", controllers.CreateVideo)
		v1.PUT("/videos/:id", controllers.UpdateVideo)
		v1.DELETE("/videos/:id", controllers.DeleteVideo)

		v1.GET("/categories", controllers.GetCategories)
		v1.GET("/categories/:id", controllers.GetCategory)
		v1.POST("/categories", controllers.CreateCategory)
		v1.PUT("categories/:id", controllers.UpdateCategory)
		v1.DELETE("/categories/:id", controllers.DeleteCategory)
		v1.GET("/categories/:id/videos", controllers.GetCategoryVideos)
	}

	r.Run()
}
