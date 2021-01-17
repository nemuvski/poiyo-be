package model

import "time"

type Comment struct {
	CommentId      string    `json:"comment_id" gorm:"primaryKey;not null;type:uuid;default:gen_random_uuid()"`
	BoardId        string    `json:"board_id" gorm:"primaryKey;not null;type:uuid"`
	OwnerAccountId string    `json:"owner_account_id" gorm:"not null;type:uuid"`
	Body           string    `json:"body" gorm:"not null;type:varchar(500)"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at"`
}
