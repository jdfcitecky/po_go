package service

import (
	. "po_go/db"
	"po_go/entity"
	"po_go/utils"

	"github.com/jinzhu/gorm"
)

type Blog entity.Blog

func (Blog) TableName() string {
	return "blog"
}

//find a blog by id
func (blog *Blog) FindOne() (b *Blog) {
	b = new(Blog)
	Db.Where("id = ?", blog.ID).First(b)
	return
}

//find the name and type by id
func (blog *Blog) FindOneTypeName() (b *Blog) {
	b = new(Blog)
	Db.Table("blog b").Select("b.*, bt.name as type_name").
		Joins("left join blog_type bt on b.typeId = bt.id").
		Where("b.id = ? ", blog.ID).Order("bt.sort asc").Find(b)
	return
}

//find the next one
func (blog *Blog) FindNextOne() (b *Blog) {
	b = new(Blog)
	result := Db.Where("id > ?", blog.ID).First(b)
	if result.Error != nil {
		return nil
	}
	return
}

//find the last one
func (blog *Blog) FindLastOne() (b *Blog) {
	b = new(Blog)
	result := Db.Where("id < ?", blog.ID).Order("id desc").First(b)
	if result.Error != nil {
		return nil
	}
	return
}

//find comment
func (blog *Blog) FindCommentByBlog() []Comment {
	comments := make([]Comment, 0)
	result := Db.Table("comment").Where("blog_id = ? and status = 1", blog.ID).Order("add_time asc").Find(&comments)
	if result.Error != nil {
		return nil
	}
	return comments
}

func (blog *Blog) FindByTypeCount() (count int) {
	Db.Model(blog).Where("typeId = ? ", blog.TypeId).Count(&count)
	return
}

//find blog list
func (blog *Blog) FindList(page *utils.Page) ([]*Blog, error) {
	bs := make([]*Blog, 0)
	curDb := Db.Table("blog b").Select("b.*, bt.name as type_name").
		Joins("left join blog_type bt on b.typeId = bt.id")
	if blog.TypeId > 0 {
		curDb = curDb.Where("b.typeId = ? ", blog.TypeId)
	}

	//Limit assign the max number of return data ; Offset assign the number of data need to be skip。
	result := curDb.Limit(page.Size).Offset(page.GetStart()).Order("`add_time` asc").Find(&bs)
	return bs, result.Error
}

func (blog *Blog) Count() (count int) {
	Db.Model(blog).Count(&count)
	return
}

func (blog *Blog) Insert() *gorm.DB {
	return Db.Create(blog)
}

func (blog *Blog) Update() *gorm.DB {
	return Db.Save(blog)
}

//update click
func (blog *Blog) UpdateClick() *gorm.DB {
	return Db.Model(blog).Where("id = ? ", blog.ID).Update("click_hit", gorm.Expr("click_hit + ?", 1))
}

func (blog *Blog) UpdateReplay() *gorm.DB {
	return Db.Model(blog).Where("id = ? ", blog.ID).Update("replay_hit", gorm.Expr("replay_hit + ?", 1))
}

func (blog *Blog) Delete() *gorm.DB {
	return Db.Delete(blog)
}
