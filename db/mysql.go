package db

import (
	"fmt"
	"po_go/conf"
	"po_go/entity"
	"po_go/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	var dbConfig = conf.Conf.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)

	Db, err = gorm.Open(conf.Conf.Db.Dialects, url)

	if err != nil {
		panic(err)
	}

	if Db.Error != nil {
		panic(Db.Error)
	}

	Db.DB().SetMaxIdleConns(dbConfig.MaxIdle)
	Db.DB().SetMaxOpenConns(dbConfig.MaxOpen)
	logger := utils.Log()
	Db.SetLogger(logger)
	Db.LogMode(true)
	//Auto migrate
	Db.AutoMigrate(&entity.Blog{}, &entity.Member{}, &entity.Work{}, &entity.Comment{}, &entity.CommentRecord{}, &entity.BrowseRecord{}, &entity.Message{})
	logger.Info("mysql connect success")

}
