package entity

type ChatRoom struct {
	ID       int  `json:"id"`
	IsActive bool `gorm:"column:is_active;defalut:true" json:"-"`
}
