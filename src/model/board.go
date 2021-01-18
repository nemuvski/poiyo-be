package model

import "time"

// Board boardsテーブルのデータモデル.
type Board struct {
	BoardId        string    `json:"board_id" gorm:"primaryKey;not null;type:uuid;default:gen_random_uuid()"`
	Title          string    `json:"title" gorm:"not null;type:varchar(200)"`
	Body           string    `json:"body" gorm:"not null;type:varchar(1000)"`
	OwnerAccountId string    `json:"owner_account_id" gorm:"not null;type:uuid"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at"`
}
