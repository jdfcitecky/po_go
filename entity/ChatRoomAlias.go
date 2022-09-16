package entity

type ChatRoomAlias struct {
	ID         int    `json:"-"`
	MemberID   int    `json:"member_id"`
	ChatRoomID int    `json:"chat_room_id"`
	Alias      string `json:"alias"`
}
