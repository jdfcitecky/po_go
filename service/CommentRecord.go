package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type CommentRecord entity.CommentRecord

func (CommentRecord) TableName() string {
	return "CommentRecords"
}

func (commentRecord *CommentRecord) FindNumberOfCommentRecordByWorkID(workID int) int {
	var count int
	Db.Model(commentRecord).
		Where("work_id = ? ", workID).
		Count(&count)

	return count
}

func (commentRecord *CommentRecord) FindNumberOfCommentRecordByWorkIDAndTime(workID int, startDate int, endDate int) int {
	var count int
	Db.Model(commentRecord).
		Where("work_id = ? ", workID).
		Where("date > ? ", startDate).
		Where("date < ? ", endDate).
		Count(&count)

	return count
}

func (commentRecord *CommentRecord) FindNumberOfCommentRecordByMemberID(memberID int) int {
	var count int
	Db.Model(commentRecord).
		Where("member_id = ? ", memberID).
		Count(&count)

	return count
}

//creat commentRecord
func (commentRecord *CommentRecord) Insert() *gorm.DB {
	return Db.Create(commentRecord)
}

func (commentRecord *CommentRecord) DeleteCommentRecord() *gorm.DB {
	return Db.Model(commentRecord).Delete(commentRecord, commentRecord.ID)
}
