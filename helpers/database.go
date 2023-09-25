package helpers

import (
	"github.com/RachidMoysePolania/MotoGuard-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("motoguard.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&models.Road_logs{}, &models.Userdata{})

	return db
}
