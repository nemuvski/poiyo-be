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

// GetBoards /boardsでボードを複数取得するAPI.
func GetBoards() echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(model.BoardsGetRequest)
		c.Bind(m)

		if err := c.Validate(m); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)

		boards := []model.Board{}
		responseQuery := tx.Model(&model.Board{}).Order("created_at DESC")

		// アカウントIDの条件を設定.
		if m.OwnerAccountId != "" {
			responseQuery = responseQuery.Where("owner_account_id = ?", m.OwnerAccountId)
		}

		// 検索キーワードの条件を設定.
		if m.Search != "" {
			responseQuery = responseQuery.Where("title LIKE ?", "%"+m.Search+"%")
		}

		// 取得範囲を設定.
		responseQuery.Offset((m.Page - 1) * m.NumPerPage).Limit(m.NumPerPage).Find(&boards)

		response := model.Boards{
			CurrentPage: m.Page,
			Items:       boards,
		}

		// 次のページにもデータがあるか判定するためのクエリを実行.
		nextBoards := []model.Board{}
		checkedNextPageResult := responseQuery.Offset(m.Page * m.NumPerPage).Limit(1).Find(&nextBoards)
		if checkedNextPageResult.RowsAffected > 0 {
			response.NextPage = m.Page + 1
		}

		return c.JSON(http.StatusOK, response)
	}
}
