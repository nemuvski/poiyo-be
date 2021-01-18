package router

import (
	"os"
	"poiyo-be/src/api"
	"poiyo-be/src/database"
	"poiyo-be/src/environment"
	customMiddleware "poiyo-be/src/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func Init() *echo.Echo {
	e := echo.New()

	if os.Getenv("GO_EXEC_ENV") == environment.EXEC_ENV_DEVELOPMENT {
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.INFO)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGIN")},
		AllowHeaders: []string{
			echo.HeaderAuthorization,
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAcceptEncoding,
		},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// リクエスト中のJWTをチェックするミドルウェアを登録.
	e.Use(customMiddleware.Auth())

	// DBサーバとの接続.
	db := database.Connect()
	// DBのトランザクションを制御するミドルウェアを登録.
	e.Use(customMiddleware.Transaction(db))

	// ルートを登録.
	v1 := e.Group("/api/v1")
	v1.POST("/auth", api.PostAccount())
	v1.POST("/board", api.PostBoard())

	return e
}
