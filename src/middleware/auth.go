package middleware

import (
	"context"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

// Auth リクエストに含まれるJWTを検証.
func Auth() echo.MiddlewareFunc {
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var opt option.ClientOption
		if os.Getenv("FB_KEYFILE_JSON") != "" {
			// (本番環境向け) 環境変数にキーファイルの内容が設定されている場合はその内容を用いる.
			opt = option.WithCredentialsJSON([]byte(os.Getenv("FB_KEYFILE_JSON")))
		} else {
			opt = option.WithCredentialsFile(os.Getenv("FB_ADMIM_SDK_KEY_PATH"))
		}
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return err
		}
		client, err := app.Auth(context.Background())
		if err != nil {
			return err
		}

		authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return err
		}

		c.Set("token", token)

		return next(c)
	}
}
