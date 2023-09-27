package helpers

import (
	"github.com/RachidMoysePolania/MotoGuard-API/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := "test:Skills39@tcp(localhost:3306)/motoguard?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&models.Userdata{}, &models.Road_logs{})

	return db
}
