package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type ChatRoomAlias entity.ChatRoomAlias

func (ChatRoomAlias) TableName() string {
	return "chat_room_aliases"
}

func (chatRoomAlias *ChatRoomAlias) FindChatRoomAliasByMemberID(memberID int) []ChatRoomAlias {
	chatRoomAliases := make([]ChatRoomAlias, 0)
	Db.Model(chatRoomAliases).
		Where("member_id = ? ", memberID).
		Find(&chatRoomAliases)

	return chatRoomAliases
}

//creat chat room alias
func (chatRoomAlias *ChatRoomAlias) Insert() *gorm.DB {
	return Db.Create(chatRoomAlias)
}

func (chatRoomAlias *ChatRoomAlias) DeleteChatRoomAlias() *gorm.DB {
	return Db.Model(chatRoomAlias).Delete(chatRoomAlias, chatRoomAlias.ID)
}
