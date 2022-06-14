package entity

import (
	"time"
)

type Comment struct {
	ID        int    `json:"id"`
	MemberID  int    `json:"member_id"`
	Member    Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Text      string `gorm:"column:txet" json:"text"`
	Status    int    `gorm:"column:status" json:"status"`
	WorkID    int    `json:"work_id"`
	Work      Work
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
