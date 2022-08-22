package entity

import (
	"time"
)

type Comment struct {
	ID         int       `json:"id"`
	MemberID   int       `json:"member_id"`
	MemberName string    `gorm:"column:member_name" json:"member_name"`
	Text       string    `gorm:"column:txet" json:"text"`
	IsNew      bool      `gorm:"column:is_new;defalut:true" json:"is_new"`
	WorkID     int       `json:"work_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
