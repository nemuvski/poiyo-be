package api

import (
	"net/http"
	customMiddleware "poiyo-be/src/middleware"
	"poiyo-be/src/model"
	"regexp"
	"strconv"
	"strings"

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

		// パスパラメータについてバリデーション.
		params := model.BoardPathParameter{Bid: boardId}
		if err := c.Validate(params); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		board := model.Board{}

		// ボードIDで検索する. (ボードIDは一意なので取得できるのは1件のみ)
		result := tx.Where("board_id = ?", boardId).First(&board)
		if result.RowsAffected == 0 {
			return c.NoContent(http.StatusNoContent)
		}
		return c.JSON(http.StatusNoContent, board)
	}
}

// GetBoards /boardsでボードを複数取得するAPI.
func GetBoards() echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParam := model.BoardsQueryParameter{
			OwnerAccountId: c.QueryParam("owner_account_id"),
			Search:         c.QueryParam("search"),
		}
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 1 {
			return c.String(http.StatusBadRequest, "pageに正の整数値を指定してください。")
		}
		queryParam.Page = page
		numPerPage, err := strconv.Atoi(c.QueryParam("num_per_page"))
		if err != nil {
			return c.String(http.StatusBadRequest, "num_per_pageに正の整数値を指定してください")
		}
		queryParam.NumPerPage = numPerPage
		if err := c.Validate(queryParam); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		boards := []model.Board{}
		responseQuery := tx.Model(&model.Board{}).Order("created_at DESC")
		// アカウントIDの条件を設定.
		if queryParam.OwnerAccountId != "" {
			responseQuery = responseQuery.Where("owner_account_id = ?", queryParam.OwnerAccountId)
		}
		// 検索キーワードの条件を設定.
		if queryParam.Search != "" {
			// 「%と_」記号がある場合はエスケープする.
			keyword := regexp.MustCompile("(%|_)").ReplaceAllString(queryParam.Search, "\\$1")
			// 空白文字をスペースに置換する.
			keyword = regexp.MustCompile("\\s").ReplaceAllString(keyword, " ")
			// キーワードの先頭と末尾の空白文字を除去.
			keyword = strings.TrimSpace(keyword)
			responseQuery = responseQuery.Where("title LIKE ?", "%"+keyword+"%")
		}
		// 範囲を設定して取得.
		responseQuery.Offset((queryParam.Page - 1) * queryParam.NumPerPage).Limit(queryParam.NumPerPage).Find(&boards)

		response := model.Boards{
			CurrentPage: queryParam.Page,
			Items:       boards,
		}

		// 次のページにもデータがあるか判定するためのクエリを実行.
		nextBoards := []model.Board{}
		checkedNextPageResult := responseQuery.Offset(queryParam.Page * queryParam.NumPerPage).Limit(1).Find(&nextBoards)
		if checkedNextPageResult.RowsAffected > 0 {
			response.NextPage = queryParam.Page + 1
		}

		return c.JSON(http.StatusOK, response)
	}
}

// DeleteBoard /boards/:bidでボードをID指定で削除するAPI.
func DeleteBoard() echo.HandlerFunc {
	return func(c echo.Context) error {
		boardId := c.Param("bid")

		// パスパラメータについてバリデーション.
		params := model.BoardPathParameter{Bid: boardId}
		if err := c.Validate(params); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		board := model.Board{}
		result := tx.Where("board_id = ?", boardId).Delete(&board)
		if result.RowsAffected == 0 {
			return c.NoContent(http.StatusNoContent)
		}
		return c.NoContent(http.StatusOK)
	}
}

// PatchBoard /boards/:bid でボードをID指定で更新するAPI.
func PatchBoard() echo.HandlerFunc {
	return func(c echo.Context) error {
		boardId := c.Param("bid")

		// パスパラメータについてバリデーション.
		params := model.BoardPathParameter{Bid: boardId}
		if err := c.Validate(params); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		m := new(model.BoardPatchRequest)
		c.Bind(m)

		if err := c.Validate(m); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		updateBoard := model.Board{
			Title: m.Title,
			Body:  m.Body,
		}
		responseBoard := model.Board{}
		// updateのインスタンスに反映結果後のレコードの内容が全てはいらない（設定したもののみ）なのでFindで反映後のレコードを取得.
		result := tx.Model(&model.Board{BoardId: boardId}).Updates(&updateBoard).Find(&responseBoard)
		if result.RowsAffected == 0 {
			return c.NoContent(http.StatusNoContent)
		}
		return c.JSON(http.StatusOK, responseBoard)
	}
}
