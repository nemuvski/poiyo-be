package api

import (
	"net/http"
	customMiddleware "poiyo-be/src/middleware"
	"poiyo-be/src/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// PostComment /commentsでコメントを作成するAPI.
func PostComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		m := new(model.CommentPostRequest)
		c.Bind(m)

		if err := c.Validate(m); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		comment := new(model.Comment)
		comment.BoardId = m.BoardId
		comment.OwnerAccountId = m.OwnerAccountId
		comment.Body = m.Body

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		tx.Create(&comment)

		return c.JSON(http.StatusCreated, comment)
	}
}

// GetComments /commentsでコメントを複数取得するAPI.
func GetComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParam := model.CommentsQueryParameter{
			BoardId: c.QueryParam("board_id"),
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
		comments := []model.Comment{}
		responseQuery := tx.Model(&model.Comment{}).Order("created_timestamp DESC")
		responseQuery = responseQuery.Where("board_id = ?", queryParam.BoardId)

		// 範囲を設定して取得.
		responseQuery.Offset((queryParam.Page - 1) * queryParam.NumPerPage).Limit(queryParam.NumPerPage).Find(&comments)

		response := model.Comments{
			CurrentPage: queryParam.Page,
			Items:       comments,
		}

		// 次のページにもデータがあるか判定するためのクエリを実行.
		nextComments := []model.Board{}
		checkedNextPageResult := responseQuery.Offset(queryParam.Page * queryParam.NumPerPage).Limit(1).Find(&nextComments)
		if checkedNextPageResult.RowsAffected > 0 {
			response.NextPage = queryParam.Page + 1
		}

		return c.JSON(http.StatusOK, response)
	}
}

// DeleteComment /comments/:bid/:cidでコメントをID指定で削除するAPI.
func DeleteComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		boardId := c.Param("bid")
		commentId := c.Param("cid")

		// パスパラメータについてバリデーション.
		params := model.DeleteCommentPathParameter{Bid: boardId, Cid: commentId}
		if err := c.Validate(params); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		comment := model.Comment{}
		result := tx.Where("comment_id = ? AND board_id = ?", commentId, boardId).Delete(&comment)
		if result.RowsAffected == 0 {
			return c.NoContent(http.StatusNoContent)
		}
		return c.NoContent(http.StatusOK)
	}
}

// PatchComment /comments/:cidでコメントをID指定で更新するAPI.
func PatchComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		commentId := c.Param("cid")

		// パスパラメータについてバリデーション.
		params := model.PatchCommentPathParameter{Cid: commentId}
		if err := c.Validate(params); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		m := new(model.CommentPatchRequest)
		c.Bind(m)

		if err := c.Validate(m); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		tx := c.Get(customMiddleware.TxKey).(*gorm.DB)
		updateComment := model.Comment{Body: m.Body}
		responseComment := model.Comment{}
		// updateのインスタンスに反映結果後のレコードの内容が全てはいらない（設定したもののみ）なのでFindで反映後のレコードを取得.
		result := tx.Model(&model.Comment{CommentId: commentId, BoardId: m.BoardId}).Updates(&updateComment).Find(&responseComment)
		if result.RowsAffected == 0 {
			return c.NoContent(http.StatusNoContent)
		}
		return c.JSON(http.StatusOK, responseComment)
	}
}
