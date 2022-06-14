package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type Comment entity.Comment

func (Comment) TableName() string {
	return "Comment"
}

func (comment *Comment) FindCommentsListByWorkID(workID int) []Comment {
	comments := make([]Comment, 0)
	Db.Table("comment c").Select("c.*, m.email as email").
		Joins("left join member m on c.member_id = m.id").
		Where("c.is_new = true").
		Order("created_at asc").
		Find(&comments)

	return comments
}

func (comment *Comment) FindCommentListAll(Map map[string]interface{}) []Comment {
	comments := make([]Comment, 0)
	Db.Model(comment).Order("created_at asc").Find(&comments)
	return comments
}

//creat comment
func (comment *Comment) Insert() *gorm.DB {
	return Db.Create(comment)
}

func (comment *Comment) UpdateStatus() *gorm.DB {
	return Db.Model(comment).Where("id = ? ", comment.ID).Update("is_new", true)
}

func (comment *Comment) DeleteComment() *gorm.DB {
	return Db.Model(comment).Delete(comment)
}
