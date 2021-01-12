package model

import "time"

type Account struct {
	AccountId   string    `json:"account_id" gorm:"primaryKey;not null;type:uuid;default:gen_random_uuid()"`
	ServiceType string    `json:"service_type" gorm:"primaryKey;not null"`
	ServiceId   string    `json:"service_id" gorm:"primaryKey;not null"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
}
