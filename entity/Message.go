package entity

import (
	"time"
)

type Message struct {
	ID         int       `json:"id"`
	MemberID   int       `json:"member_id"`
	ReceiverID int       `json:"receiver_id"`
	Date       int       `json:"date"`
	Text       string    `json:"text"`
	IsHide     bool      `gorm:"column:is_hide;defalut:false" json:"is_hide"`
	IsRead     bool      `gorm:"column:is_read;defalut:false" json:"is_read"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
