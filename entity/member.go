package entity

import (
	"time"
)

type Member struct {
	ID        int       `json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	IsManager int       `gorm:"column:is_manager" json:"is_manager"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
