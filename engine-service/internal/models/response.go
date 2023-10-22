package models

type Response struct {
	Data []int   `json:"data"`
	Err  *string `json:"err"`
}
