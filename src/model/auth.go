package model

// AuthPostRequest /authのPOSTのリクエストボディ.
type AuthPostRequest struct {
	ServiceType string `json:"service_type"`
	ServiceId   string `json:"service_id"`
	Email       string `json:"email"`
}
