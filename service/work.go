package service

import (
	"fmt"
	. "po_go/db"
	"po_go/entity"
	"po_go/utils"

	"github.com/jinzhu/gorm"
)

type Work entity.Work

func (Work) TableName() string {
	return "works"
}

func (work *Work) FindOne() (w *Work) {
	w = new(Work)
	Db.Model(work).Where("id = ?", work.ID).First(w)
	return
}

func (work *Work) FindAll() []*Work {
	ws := make([]*Work, 0)
	Db.Model(work).Order("created_at asc").Find(&ws)
	return ws
}
func (work *Work) Search(keyword string) []*Work {
	ws := make([]*Work, 0)
	Db.Model(work).Where(fmt.Sprintf(" tags like %q", ("%" + keyword + "%"))).Order("created_at asc").Find(&ws)
	return ws
}

func (work *Work) SearchTool(keyword string) []*Work {
	ws := make([]*Work, 0)
	Db.Model(work).Where(fmt.Sprintf(" tools like %q", ("%" + keyword + "%"))).Order("created_at asc").Find(&ws)
	return ws
}
func (work *Work) SearchCategory(keyword string) []*Work {
	ws := make([]*Work, 0)
	Db.Model(work).Where(fmt.Sprintf(" category like %q", ("%" + keyword + "%"))).Order("created_at asc").Find(&ws)
	return ws
}

//update click
func (work *Work) UpdateClick() *gorm.DB {
	return Db.Model(work).Where("id = ? ", work.ID).Update("click_hit", gorm.Expr("click_hit + ?", 1))
}

func (work *Work) Count() (count int) {
	Db.Model(work).Count(&count)
	return
}

func (work *Work) Insert() *gorm.DB {
	logger := utils.Log()
	logger.Info("work insert")
	return Db.Create(work)
}

func (work *Work) Update() *gorm.DB {
	return Db.Save(work)
}

func (work *Work) Delete() *gorm.DB {
	return Db.Delete(work)
}
