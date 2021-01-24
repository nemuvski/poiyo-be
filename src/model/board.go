package model

import "time"

// BoardPostRequest /boardsのPOSTのリクエストボディ.
type BoardPostRequest struct {
	Title          string `json:"title" validate:"required,max=200"`
	Body           string `json:"body" validate:"required,max=1000"`
	OwnerAccountId string `json:"owner_account_id" validate:"required,uuid4"`
}

// BoardsGetRequest /boardsのGETのリクエストボディ.
type BoardsGetRequest struct {
	Page           int    `json:"page" validate:"required,min=1"`
	NumPerPage     int    `json:"num_per_page" validate:"required,min=1,max=50"`
	OwnerAccountId string `json:"owner_account_id" validate:"omitempty,uuid4"`
	Search         string `json:"search" validate:"omitempty"`
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