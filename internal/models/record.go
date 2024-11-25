package models

import "time"

type URLMapping struct {
	ID          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}
