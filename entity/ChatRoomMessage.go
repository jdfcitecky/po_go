package entity

type ChatRoomMessage struct {
	ID         int    `json:"-"`
	SenderID   int    `json:"sender_id"`
	ChatRoomID int    `json:"chat_room_id"`
	Date       int    `json:"date"`
	Time       string `json:"time"`
	Text       string `json:"text"`
	IsHide     bool   `gorm:"column:is_hide;defalut:false" json:"is_hide"`
	IsRead     bool   `gorm:"column:is_read;defalut:false" json:"is_read"`
}
