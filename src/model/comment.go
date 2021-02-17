package model

import (
	"database/sql"
	"time"
)

// CommentPostRequest /commentsのPOSTのリクエストボディ.
type CommentPostRequest struct {
	BoardId        string `json:"board_id" validate:"required,uuid4"`
	OwnerAccountId string `json:"owner_account_id" validate:"required,uuid4"`
	Body           string `json:"body" validate:"required,max=500"`
}

// CommentPatchRequest /commentsのPATCHのリクエストボディ.
type CommentPatchRequest struct {
	BoardId string `json:"board_id" validate:"required,uuid4"`
	Body    string `json:"body" validate:"required,max=500"`
}

// CommentsQueryParameter /commentsのGETのクエリパラメータ.
type CommentsQueryParameter struct {
	Page       int    `validate:"required,min=1"`
	NumPerPage int    `validate:"required,min=1,max=50"`
	BoardId    string `validate:"required,uuid4"`
}

// DeleteCommentPathParameter /comments/:bid/:cidのパスパラメータ.
type DeleteCommentPathParameter struct {
	Bid string `validate:"required,uuid4"`
	Cid string `validate:"required,uuid4"`
}

// PatchCommentPathParameter /comments/:cidのパスパラメータ.
type PatchCommentPathParameter struct {
	Cid string `validate:"required,uuid4"`
}

// Comment commentsテーブルのデータモデル.
type Comment struct {
	CommentId        string       `json:"comment_id" gorm:"primaryKey;not null;type:uuid;default:gen_random_uuid()"`
	BoardId          string       `json:"board_id" gorm:"primaryKey;not null;type:uuid"`
	OwnerAccountId   string       `json:"owner_account_id" gorm:"not null;type:uuid"`
	Body             string       `json:"body" gorm:"not null;type:varchar(500)"`
	CreatedTimestamp time.Time    `json:"created_timestamp" gorm:"not null;default:current_timestamp"`
	UpdatedTimestamp sql.NullTime `json:"updated_timestamp" gorm:"default:null"`
}

type Comments struct {
	NextPage    int       `json:"next_page,omitempty"`
	CurrentPage int       `json:"current_page"`
	Items       []Comment `json:"items"`
}
