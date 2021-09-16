package models

type MResponse struct {
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
