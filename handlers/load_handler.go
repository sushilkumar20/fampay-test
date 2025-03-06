package handlers

import (
	"fam/db"
	"fam/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var videQueryHandler *VideoQueryHandler

func LoadHandler(router *gin.Engine, database *gorm.DB) {
	videQueryHandler = loadVideoQueryHandler(database)
	loadRoutes(router)
}

func loadVideoQueryHandler(database *gorm.DB) *VideoQueryHandler {
	dbClient := db.NewPostgresStore(database)
	videoQueryService := service.NewVideoQueryServiceImpl(dbClient)
	return NewVideoQueryHandler(videoQueryService)
}

func loadRoutes(router *gin.Engine) {
	//authenticate
	user := router.Group("/api/v1")

	user.Use()
	{
		user.GET("/videos", videQueryHandler.GetVideos)
		user.GET("/videos/search", videQueryHandler.SearchVideos)
	}
}
