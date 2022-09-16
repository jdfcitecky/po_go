package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type ChatRoomMember entity.ChatRoomMember

func (ChatRoomMember) TableName() string {
	return "chat_room_members"
}

func (chatRoomMember *ChatRoomMember) FindChatRoomMemberListByMemberID(chatRoomID int) []ChatRoomMember {
	chatRoomMembers := make([]ChatRoomMember, 0)
	Db.Model(chatRoomMember).
		Where("chat_room_id = ? ", chatRoomID).
		Find(&chatRoomMembers)

	return chatRoomMembers
}

//creat message
func (chatRoomMember *ChatRoomMember) Insert() *gorm.DB {
	return Db.Create(chatRoomMember)
}

func (chatRoomMember *ChatRoomMember) DeleteChatRoomMember() *gorm.DB {
	return Db.Model(chatRoomMember).Delete(chatRoomMember, chatRoomMember.ID)
}
