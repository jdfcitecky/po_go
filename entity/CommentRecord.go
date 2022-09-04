package entity

import (
	"time"
)

type CommentRecord struct {
	ID        int       `json:"id"`
	MemberID  int       `json:"member_id"`
	WorkID    int       `json:"work_id"`
	IP        string    `json:"ip"`
	Date      int       `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
