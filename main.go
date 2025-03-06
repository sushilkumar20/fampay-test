package main

import (
	"fam/config"
	"fam/db"
	"fam/handlers"
	"fam/probe"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	//Load database Schema and Cron job
	database := db.InitDb(config.Configuration.PostgresConfig)
	db.LoadSchema(database)
	LoadProbeJob(database)

	handlers.LoadHandler(router, database)
	public := router.Group("/api/v1")
	public.GET("/health", Health)
	err := router.Run("0.0.0.0:" + strconv.Itoa(config.Configuration.ServerPort))

	if err != nil {
		log.Panic("Failed to start the server, Reason: ", err)
		return
	}

	log.Println("Server started")
}

func Health(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEPlain, []byte(`0`))
}

func LoadProbeJob(dataBase *gorm.DB) {
	probe.LoadProbe(dataBase)
}
