package db

import "gorm.io/gorm"

func LoadSchema(database *gorm.DB) error {
	err := database.AutoMigrate(&VideoMetaData{})
	if err != nil {
		return err
	}
	return nil
}
