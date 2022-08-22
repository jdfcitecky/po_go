package entity

import (
	"time"
)

type Member struct {
	ID        int       `json:"id"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	IsManager bool      `gorm:"column:is_manager;default:false" json:"is_manager"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
