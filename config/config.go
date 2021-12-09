package config

import (
	"rizadwi.com/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	var err error
	dsn := "host=localhost user=postgres password=riza4231 dbname=training port=5432 sslmode=disable timezone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Article{})
	if err != nil {
		panic(err)
	}
}
