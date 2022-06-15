package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type Member entity.Member

func (Member) TableName() string {
	return "members"
}

// find a member by username
func (member *Member) Login() (m *Member) {
	m = new(Member)
	Db.Model(member).Where("email = ? ", member.Email).First(m)
	return
}

// find a member's name
func (member *Member) FindMemberByID(memberID int) (m *Member) {
	m = new(Member)
	Db.Model(member).Where("id = ? ", memberID).First(m)
	return
}

// find a manager
func (member *Member) Find() (m *Member) {
	m = new(Member)
	Db.Model(member).Where("IsManager = true").First(m)
	return
}

// create a member
func (member *Member) Insert() *gorm.DB {
	return Db.Create(member)
}

//modify member info
func (member *Member) UpdateInfo() *gorm.DB {
	if member.Password != "" {
		return Db.Model(member).Update(member)
	}
	return Db.Save(member)
}

//modify member password
func (member *Member) UpdatePassword() *gorm.DB {
	if member.Password != "" {
		return Db.Model(member).Select("password").Update(member)
	}
	return Db.Save(member)
}
