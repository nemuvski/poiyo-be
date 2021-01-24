package api

import (
	"net/http"
	customMiddleware "poiyo-be/src/middleware"
	"poiyo-be/src/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// PostBoard /boardsでボードを作成するAPI.
func PostBoard() echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(model.BoardPostRequest)
		c.Bind(m)

		if err := c.Validate(m); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		board := new(model.Board)
		board.Title = m.Title
		board.Body = m.Body
		board.OwnerAccountId = m.OwnerAccountId

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		tx.Create(&board)

		return c.JSON(http.StatusCreated, board)
	}
}

// GetBoard /boards/:bidでボードをID指定で取得するAPI.
func GetBoard() echo.HandlerFunc {
	return func(c echo.Context) error {
		boardId := c.Param("bid")

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		board := model.Board{}

		// ボードIDで検索する. (ボードIDは一意なので取得できるのは1件のみ)
		result := tx.Where("board_id = ?", boardId).First(&board)

		// 取得できたか否かで、ステータスコードを変える.
		status := http.StatusOK
		if result.RowsAffected == 0 {
			status = http.StatusNoContent
		}

		return c.JSON(status, board)
	}
}
