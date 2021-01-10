package database

import (
	"fmt"
	"log"
	"poiyo-be/src/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dotenv *environment.DotEnv) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dotenv.Host,
		dotenv.User,
		dotenv.Pass,
		dotenv.Db,
		dotenv.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("接続失敗", err)
	}
	return db
}
