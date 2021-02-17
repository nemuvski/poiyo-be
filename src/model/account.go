package model

import "time"

// DeleteAccountPathParameter /accounts/:aidのパスパラメータ.
type DeleteAccountPathParameter struct {
	Aid string `validate:"required,uuid4"`
}

// Account accountsテーブルのデータモデル.
type Account struct {
	AccountId        string    `json:"account_id" gorm:"primaryKey;not null;type:uuid;default:gen_random_uuid()"`
	ServiceType      string    `json:"service_type" gorm:"primaryKey;not null"`
	ServiceId        string    `json:"service_id" gorm:"primaryKey;not null"`
	Email            string    `json:"email"`
	CreatedTimestamp time.Time `json:"created_timestamp" gorm:"not null;default:current_timestamp"`
}
