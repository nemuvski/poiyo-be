package model

import "time"

// BoardPostRequest /boardsのPOSTのリクエストボディ.
type BoardPostRequest struct {
	Title          string `json:"title" validate:"required,max=200"`
	Body           string `json:"body" validate:"required,max=1000"`
	OwnerAccountId string `json:"owner_account_id" validate:"required,uuid4"`
}

// BoardPatchRequest /boardsのPATCHのリクエストボディ.
type BoardPatchRequest struct {
	Title string `json:"title" validate:"required,max=200"`
	Body  string `json:"body" validate:"required,max=1000"`
}

// BoardsQueryParameter /boardsのGETのクエリパラメータ.
type BoardsQueryParameter struct {
	Page           int    `validate:"required,min=1"`
	NumPerPage     int    `validate:"required,min=1,max=50"`
	OwnerAccountId string `validate:"omitempty,uuid4"`
	Search         string `validate:"omitempty"`
}

// BoardPathParameter 削除・編集・単一取得のAPIのパスパラメータ.
type BoardPathParameter struct {
	Bid string `validate:"required,uuid4"`
}

// Board boardsテーブルのデータモデル.
type Board struct {
	BoardId        string    `json:"board_id" gorm:"primaryKey;not null;type:uuid;default:gen_random_uuid()"`
	Title          string    `json:"title" gorm:"not null;type:varchar(200)"`
	Body           string    `json:"body" gorm:"not null;type:varchar(1000)"`
	OwnerAccountId string    `json:"owner_account_id" gorm:"not null;type:uuid"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Boards struct {
	NextPage    int     `json:"next_page,omitempty"`
	CurrentPage int     `json:"current_page"`
	Items       []Board `json:"items"`
}
