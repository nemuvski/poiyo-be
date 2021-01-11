package main

import (
	"fmt"
	"net/http"
	"poiyo-be/src/database"
	"poiyo-be/src/environment"
	"poiyo-be/src/model"

	"github.com/labstack/echo"
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

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
