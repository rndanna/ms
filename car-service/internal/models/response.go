package models

type APIResponse struct {
	Data Engine  `json:"data"`
	Err  *string `json:"err"`
}
