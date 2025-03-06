package db

import (
	"gorm.io/gorm"
	"log"
)

func LoadSchema(database *gorm.DB) {
	err := database.AutoMigrate(&VideoMetaData{})
	if err != nil {
		log.Panic("failed to migrate schema", err)
	}
}
