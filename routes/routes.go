package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/olooeez/video-vault/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		videos := v1.Group("/videos")
		{
			videos.GET("", controllers.GetVideos)
			videos.GET("/:id", controllers.GetVideo)
			videos.POST("", controllers.CreateVideo)
			videos.PUT("/:id", controllers.UpdateVideo)
			videos.DELETE("/:id", controllers.DeleteVideo)
			videos.GET("/search", controllers.SearchVideos)
		}

		categories := v1.Group("/categories")
		{
			categories.GET("", controllers.GetCategories)
			categories.GET("/:id", controllers.GetCategory)
			categories.POST("", controllers.CreateCategory)
			categories.PUT("/:id", controllers.UpdateCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
			categories.GET("/:id/videos", controllers.GetCategoryVideos)
		}
	}

	return r
}

func StartServer() {
	r := SetupRouter()
	r.Run()
}
