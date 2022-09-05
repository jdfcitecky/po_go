package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type BrowseRecord entity.BrowseRecord

type TopBrowseRecord struct {
	WorkID int    `json:"work_id"`
	Count  int    `json:"count"`
	Title  string `json:"title"`
}

type DailyBrowseRecord struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func (BrowseRecord) TableName() string {
	return "browse_records"
}

func (browseRecord *BrowseRecord) FindNumberOfBrowseRecordByWorkID(workID int) int {
	var count int
	Db.Table("browse_records").
		Where("work_id = ? ", workID).
		Count(&count)

	return count
}

func (browseRecord *BrowseRecord) FindNumberOfBrowseRecordByIP(ip string) int {
	var count int
	Db.Table("browse_records").
		Where("ip = ? ", ip).
		Count(&count)

	return count
}

func (browseRecord *BrowseRecord) FindNumberOfBrowseRecordGroupByWorkID() []TopBrowseRecord {
	result := make([]TopBrowseRecord, 0)
	Db.Table("browse_records").
		Select("work_id as work_id,COUNT(*) as count,works.title as title").
		Group("work_id").
		Joins("left join works on works.id = browse_records.work_id").
		Order("count desc").
		Scan(&result)

	return result
}

func (browseRecord *BrowseRecord) FindNumberOfBrowseRecordGroupByDate() []DailyBrowseRecord {
	result := make([]DailyBrowseRecord, 0)
	Db.Table("browse_records").
		Select("date as date,COUNT(*) as count").
		Group("date").
		Order("date").
		Scan(&result)

	return result
}

func (browseRecord *BrowseRecord) FindNumberOfBrowseRecordByWorkIDAndTime(workID int, startDate int, endDate int) int {
	var count int
	Db.Table("browse_records").
		Where("work_id = ? ", workID).
		Where("date > ? ", startDate).
		Where("date < ? ", endDate).
		Count(&count)

	return count
}

//creat browseRecord
func (browseRecord *BrowseRecord) Insert() *gorm.DB {
	return Db.Create(browseRecord)
}

func (browseRecord *BrowseRecord) DeleteBrowseRecord() *gorm.DB {
	return Db.Model(browseRecord).Delete(browseRecord, browseRecord.ID)
}
