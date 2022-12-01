package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type Message entity.Message

func (Message) TableName() string {
	return "messages"
}

func (message *Message) FindMessagesListByMemberID(memberID int) []Message {
	messages := make([]Message, 0)
	Db.Model(message).
		Where("member_id = ? ", memberID).
		Where("is_hide = false").
		Order("created_at asc").
		Find(&messages)

	return messages
}

func (message *Message) FindNumberOfMessagesByMemberIDAndReceiverID(memberID int, receiverID int) []Message {
	messages := make([]Message, 0)
	Db.Model(message).
		Where("member_id = ? ", memberID).
		Where("receiver_id = ? ", receiverID).
		Where("is_hide = false").
		Order("created_at asc").
		Find(&messages)

	return messages
}

//creat message
func (message *Message) Insert() *gorm.DB {
	return Db.Create(message)
}

//For hide message
func (message *Message) UpdateStatusHide() *gorm.DB {
	return Db.Model(message).Where("id = ? ", message.ID).Update("is_hide", true)
}

//For read message
func (message *Message) UpdateStatusRead() *gorm.DB {
	return Db.Model(message).Where("id = ? ", message.ID).Update("is_read", true)
}

func (message *Message) DeleteMessage() *gorm.DB {
	return Db.Model(message).Delete(message, message.ID)
}
