package model

// AuthPostRequest /authのPOSTのリクエストボディ.
type AuthPostRequest struct {
	ServiceType string `json:"service_type" validate:"required"`
	ServiceId   string `json:"service_id" validate:"required"`
	Email       string `json:"email,omitempty"`
}
