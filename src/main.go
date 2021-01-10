package main

import (
	"fmt"
	"poiyo-be/src/database"
	"poiyo-be/src/environment"
	"poiyo-be/src/model"
)

func main() {
	dotenv := environment.Load()
	db := database.Connect(dotenv)

	// usersテーブルから値を取得して出力.
	users := []model.User{}
	db.Find(&users)
	for _, user := range users {
		fmt.Printf("%s %s\n", user.UserId, user.CreatedAt)
	}
}
