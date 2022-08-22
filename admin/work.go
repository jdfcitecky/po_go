package admin

import (
	"po_go/service"
	"po_go/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//find work list for manage panel
func WorkList(c *gin.Context) {
	json := make(map[string]interface{})
	var work service.Work
	err := c.ShouldBind(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}

	result := work.FindAll()

	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)

}

//Find a work for manage panel
func FindWorkForManage(c *gin.Context) {
	var work service.Work
	//bind work id
	err := c.BindJSON(&work)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	//find work by id
	result := work.FindOne()

	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

//Save work
func WorkSave(c *gin.Context) {
	var work service.Work
	err := c.BindJSON(&work)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}

	var result *gorm.DB

	if work.ID <= 0 {
		result = work.Insert()
	} else {
		result = work.Update()
	}

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "Save error"}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)

}

//Delete work
func WorkDelete(c *gin.Context) {
	var work service.Work
	err := c.BindJSON(&work)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format error"}
		res.Json(c)
		return
	}

	result := work.Delete()

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "Delete error"}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}
