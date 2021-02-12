package router

import (
	"os"
	"poiyo-be/src/api"
	"poiyo-be/src/database"
	"poiyo-be/src/environment"
	customMiddleware "poiyo-be/src/middleware"
	"poiyo-be/src/validation"

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
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))

	// リクエスト中のJWTをチェックするミドルウェアを登録.
	e.Use(customMiddleware.Auth())

	// DBサーバとの接続.
	db := database.Connect()
	// DBのトランザクションを制御するミドルウェアを登録.
	e.Use(customMiddleware.Transaction(db))

	// バリデーションを登録.
	e.Validator = validation.NewValidator()

	v1 := e.Group("/api/v1")
	// 認証関連.
	v1.POST("/auth", api.PostAccount())
	// アカウント関連.
	v1.DELETE("/accounts/:aid", api.DeleteAccount())
	// ボード関連.
	v1.POST("/boards", api.PostBoard())
	v1.GET("/boards", api.GetBoards())
	v1.GET("/boards/:bid", api.GetBoard())
	v1.DELETE("/boards/:bid", api.DeleteBoard())
	v1.PATCH("/boards/:bid", api.PatchBoard())
	// コメント関連.
	v1.POST("/comments", api.PostComment())
	v1.GET("/comments", api.GetComments())
	v1.DELETE("/comments/:bid/:cid", api.DeleteComment())
	v1.PATCH("/comments/:cid", api.PatchComment())

	return e
}
