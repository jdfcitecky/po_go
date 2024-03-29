package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type ChatRoomMessage entity.ChatRoomMessage

func (ChatRoomMessage) TableName() string {
	return "chat_room_messages"
}

func (chatRoomMessage *ChatRoomMessage) FindChatRoomMessagesListByChatRoomID(chatRoomID int) []ChatRoomMessage {
	chatRoomMessages := make([]ChatRoomMessage, 0)
	Db.Model(chatRoomMessage).
		Where("chat_room_id = ? ", chatRoomID).
		Where("is_hide = false").
		Find(&chatRoomMessages)

	return chatRoomMessages
}

func (chatRoomMessage *ChatRoomMessage) FindChatRoomMessagesListByChatRoomIDWithLimit(chatRoomID int, pageStart int, pageLimit int) []ChatRoomMessage {
	chatRoomMessages := make([]ChatRoomMessage, 0)
	Db.Model(chatRoomMessage).
		Where("chat_room_id = ?", chatRoomID).
		Where("is_hide = false").
		Limit(pageLimit).
		Offset(pageStart).
		Order("id desc").
		Find(&chatRoomMessages)
	// swap to make the data sort by time asc
	for i, j := 0, len(chatRoomMessages)-1; i < j; i, j = i+1, j-1 {
		chatRoomMessages[i], chatRoomMessages[j] = chatRoomMessages[j], chatRoomMessages[i]
	}

	return chatRoomMessages
}

func (chatRoomMessage *ChatRoomMessage) UpdateMultiStatusRead() *gorm.DB {
	return Db.Model(chatRoomMessage).Where("sender_id <> ? ", chatRoomMessage.SenderID).Where("chat_room_id = ? ", chatRoomMessage.ChatRoomID).Update("is_read", true)
}

//creat message
func (chatRoomMessage *ChatRoomMessage) Insert() *gorm.DB {
	return Db.Create(chatRoomMessage)
}

//For hide message
func (chatRoomMessage *ChatRoomMessage) UpdateStatusHide() *gorm.DB {
	return Db.Model(chatRoomMessage).Where("id = ? ", chatRoomMessage.ID).Update("is_hide", true)
}

//For read message
func (chatRoomMessage *ChatRoomMessage) UpdateStatusRead() *gorm.DB {
	return Db.Model(chatRoomMessage).Where("id = ? ", chatRoomMessage.ID).Update("is_read", true)
}

func (chatRoomMessage *ChatRoomMessage) DeleteMessage() *gorm.DB {
	return Db.Model(chatRoomMessage).Delete(chatRoomMessage, chatRoomMessage.ID)
}
