package service

import (
	"fmt"
	. "po_go/db"
	"po_go/entity"

	"github.com/jinzhu/gorm"
)

type Work entity.Work

func (Work) TableName() string {
	return "blog_type"
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
	Db.Model(work).Where(fmt.Sprintf(" tools like %q", ("%" + keyword + "%"))).Order("created_at asc").Find(&ws)
	return ws
}

func (work *Work) Count() (count int) {
	Db.Model(work).Count(&count)
	return
}

func (work *Work) Insert() *gorm.DB {
	return Db.Create(work)
}

func (work *Work) Update() *gorm.DB {
	return Db.Save(work)
}

func (work *Work) Delete() *gorm.DB {
	return Db.Delete(work)
}
