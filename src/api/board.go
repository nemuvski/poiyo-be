package api

import (
	"net/http"
	customMiddleware "poiyo-be/src/middleware"
	"poiyo-be/src/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// PostBoard /boardでボードを作成するAPI.
func PostBoard() echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(model.BoardPostRequest)
		c.Bind(&m)

		board := new(model.Board)
		board.Title = m.Title
		board.Body = m.Body
		board.OwnerAccountId = m.OwnerAccountId

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		tx.Create(&board)

		return c.JSON(http.StatusCreated, board)
	}
}
