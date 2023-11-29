package database

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Friendship{}, &models.Like{}); err != nil {
		panic(err)
	}
}
