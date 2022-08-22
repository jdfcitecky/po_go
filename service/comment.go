package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type Comment entity.Comment

func (Comment) TableName() string {
	return "Comments"
}

func (comment *Comment) FindCommentsListByWorkID(workID int) []Comment {
	comments := make([]Comment, 0)
	Db.Model("comment").
		Where("work_id = ? ", workID).
		Where("is_new = false").
		Order("created_at asc").
		Find(&comments)

	return comments
}

func (comment *Comment) FindCommentListAll() []Comment {
	comments := make([]Comment, 0)
	Db.Model(comment).Order("created_at asc").Find(&comments)
	return comments
}

//creat comment
func (comment *Comment) Insert() *gorm.DB {
	return Db.Create(comment)
}

//For review comment
func (comment *Comment) UpdateStatus() *gorm.DB {
	return Db.Model(comment).Where("id = ? ", comment.ID).Update("is_new", false)
}

func (comment *Comment) DeleteComment() *gorm.DB {
	return Db.Model(comment).Delete(comment)
}
