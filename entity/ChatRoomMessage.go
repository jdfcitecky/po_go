package entity

type ChatRoomMessage struct {
	ID         int    `json:"-"`
	SenderID   int    `json:"sender_id"`
	ChatRoomID int    `json:"-"`
	Date       int    `json:"date"`
	Text       string `json:"text"`
	IsHide     bool   `gorm:"column:is_hide;defalut:false" json:"-"`
	IsRead     bool   `gorm:"column:is_read;defalut:false" json:"is_read"`
}
