package entity

import (
	"time"
)

type Blog struct {
	ID        uint      `json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	TypeId    int       `gorm:"column:typeId" json:"typeId"`
	Content   string    `gorm:"column:content" json:"content"`
	Summary   string    `gorm:"column:summary" json:"summary"`
	ClickHit  int       `gorm:"column:click_hit" json:"click_hit"`
	ReplayHit int       `gorm:"column:replay_hit" json:"replay_hit"`
	AddTime   string    `gorm:"column:add_time" json:"add_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	TypeName  string    `gorm:"-" json:"type_name"`
}
