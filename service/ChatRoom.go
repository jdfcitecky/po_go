package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type ChatRoom entity.ChatRoom

func (ChatRoom) TableName() string {
	return "chat_rooms"
}

//creat chatRoom
func (chatRoom *ChatRoom) Insert() *gorm.DB {
	return Db.Create(chatRoom)
}

func (chatRoom *ChatRoom) DeleteChatRoom() *gorm.DB {
	return Db.Model(chatRoom).Delete(chatRoom, chatRoom.ID)
}
