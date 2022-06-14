package entity

import (
	"time"
)

type Work struct {
	ID           int       `json:"id"`
	Category     string    `gorm:"column:category" json:"category"`
	Title        string    `gorm:"column:title" json:"title"`
	Text         string    `gorm:"column:text" json:"text"`
	Tools        string    `gorm:"column:tools" json:"tools"`
	Year         string    `gorm:"column:year" json:"year"`
	Downloadlink string    `gorm:"column:downloadlink" json:"downloadlink"`
	Pictureone   string    `gorm:"column:pictureone" json:"pictureone"`
	Picturetwo   string    `gorm:"column:picturetwo" json:"picturetwo"`
	Picturethree string    `gorm:"column:picturethree" json:"picturethree"`
	Picturefour  string    `gorm:"column:picturefour" json:"picturefour"`
	Picturefive  string    `gorm:"column:picturefive" json:"picturefive"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
