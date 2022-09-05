package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type CommentRecord entity.CommentRecord

type TopCommentRecord struct {
	WorkID int    `json:"work_id"`
	Count  int    `json:"count"`
	Title  string `json:"title"`
}

type DailyCommentRecord struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func (CommentRecord) TableName() string {
	return "comment_records"
}

func (commentRecord *CommentRecord) FindNumberOfCommentRecordByWorkID(workID int) int {
	var count int
	Db.Model(commentRecord).
		Where("work_id = ? ", workID).
		Count(&count)

	return count
}

func (commentRecord *CommentRecord) FindNumberOfCommentRecordByIP(ip string) int {
	var count int
	Db.Model(commentRecord).
		Where("ip = ? ", ip).
		Count(&count)

	return count
}

func (commentRecord *CommentRecord) FindNumberOfCommentRecordGroupByWorkID() []TopCommentRecord {
	result := make([]TopCommentRecord, 0)
	Db.Table("comment_records").
		Select("work_id as work_id,COUNT(*) as count,works.title as title").
		Group("work_id").
		Joins("left join works on works.id = comment_records.work_id").
		Find(&result)

	return result
}

func (commentRecord *CommentRecord) FindNumberOfCommentRecordGroupByDate() []DailyCommentRecord {
	result := make([]DailyCommentRecord, 0)
	Db.Table("comment_records").
		Select("date as date,COUNT(*) as count").
		Group("date").
		Find(&result)

	return result
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
