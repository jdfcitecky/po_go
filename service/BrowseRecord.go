package service

import (
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type BrowseRecord entity.BrowseRecord

func (BrowseRecord) TableName() string {
	return "browse_records"
}

func (browseRecord *BrowseRecord) FindNumberOfBrowseRecordByWorkID(workID int) int {
	var count int
	Db.Model(browseRecord).
		Where("work_id = ? ", workID).
		Count(&count)

	return count
}

func (browseRecord *BrowseRecord) FindNumberOfBrowseRecordByWorkIDAndTime(workID int, startDate int, endDate int) int {
	var count int
	Db.Model(browseRecord).
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
