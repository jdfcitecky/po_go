package entity

import (
	"time"
)

type BrowseRecord struct {
	ID        int       `json:"id"`
	WorkID    int       `json:"work_id"`
	IP        string    `json:"ip"`
	Date      int       `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
