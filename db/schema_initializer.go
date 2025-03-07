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

	// Create a GIN index for full-text search on title and description
	database.Exec("CREATE INDEX IF NOT EXISTS idx_title_description ON videos USING GIN (to_tsvector('english', title || ' ' || description));")

}
