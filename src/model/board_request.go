package model

// BoardPostRequest /boardのPOSTのリクエストボディ.
type BoardPostRequest struct {
	Title          string `json:"title"`
	Body           string `json:"body"`
	OwnerAccountId string `json:"owner_account_id"`
}
