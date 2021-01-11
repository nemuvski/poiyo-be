package router

import (
	"net/http"
	"os"
	"poiyo-be/src/environment"
	customMiddleware "poiyo-be/src/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func Init() *echo.Echo {
	e := echo.New()

	// 実行モードによってログの出力レベルを変更.
	if os.Getenv("GO_EXEC_ENV") == environment.EXEC_ENV_DEVELOPMENT {
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.INFO)
	}

	// ミドルウェアの登録.
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
		AllowMethods: []string{echo.GET},
	}))
	e.Use(customMiddleware.Auth())

	// ルートを登録.
	v1 := e.Group("/api/v1")
	v1.GET("/auth", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
