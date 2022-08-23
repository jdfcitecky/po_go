package entity

import (
	"time"
)

type Comment struct {
	ID         int       `json:"id"`
	MemberID   int       `json:"member_id"`
	MemberName string    `gorm:"column:member_name" json:"member_name"`
	Member     Member    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Text       string    `gorm:"column:txet" json:"text"`
	IsNew      bool      `gorm:"column:is_new;defalut:true" json:"is_new"`
	WorkID     int       `json:"work_id"`
	WorkName   string    `gorm:"column:work_name" json:"work_name"`
	Work       Work      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
