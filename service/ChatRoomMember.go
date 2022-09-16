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

//creat chat room member
func (chatRoomMember *ChatRoomMember) Insert(memberID int, chatRoomID int) *gorm.DB {
	newChatRoomMember := ChatRoomMember{MemberID: memberID, ChatRoomID: chatRoomID}
	return Db.Create(&newChatRoomMember)
}

func (chatRoomMember *ChatRoomMember) DeleteChatRoomMember() *gorm.DB {
	return Db.Model(chatRoomMember).Delete(chatRoomMember, chatRoomMember.ID)
}
