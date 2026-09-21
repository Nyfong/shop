package model

import "time"

type APIResponse struct {
	Status   int         `json:"status"`
	DateTime time.Time   `json:"datetime"`
	Payload  interface{} `json:"payload"`
}