package database

import (
	"fmt"
	"log"
	"os"
	"poiyo-be/src/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect DBサーバーに接続.
func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if os.Getenv("GO_EXEC_ENV") == environment.EXEC_ENV_DEVELOPMENT {
		db.Logger.LogMode(logger.Info)
	} else {
		db.Logger.LogMode(logger.Error)
	}

	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	return db
}
