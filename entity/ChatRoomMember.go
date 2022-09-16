package entity

type ChatRoomMember struct {
	ID         int `json:"-"`
	MemberID   int `json:"member_id"`
	ChatRoomID int `json:"chat_room_id"`
}
