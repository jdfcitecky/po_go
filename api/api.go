package api

import (
	"fmt"
	"po_go/service"
	"po_go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//find manager
func FindManager(c *gin.Context) {
	var member service.Member
	result := member.Find()
	result.Password = ""
	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

//Find the comment list for a work
func CommentListForWork(c *gin.Context) {
	json := make(map[string]interface{})
	err := c.ShouldBind(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	comment := new(service.Comment)

	//get work Id
	workID, err := strconv.Atoi(utils.StrVal(json["work_id"]))
	if err == nil {
		comment.WorkID = workID
	}
	//query comment list
	result := comment.FindCommentsListByWorkID(workID)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: err.Error()}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

//Find a work
func FindWork(c *gin.Context) {
	var work service.Work
	var comment service.Comment
	//bind work id
	err := c.BindJSON(&work)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	//update the number of click times
	work.UpdateClick()

	//find work by id
	result := work.FindOne()

	//find comment
	comments := comment.FindCommentsListByWorkID(work.ID)
	Map := make(map[string]interface{})
	Map["blog"] = result
	Map["comments"] = comments
	res := &utils.Response{Code: 0, Msg: "", Data: Map}
	res.Json(c)
}

//Find the work list
func Search(c *gin.Context) {
	json := make(map[string]interface{})
	err := c.ShouldBind(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	work := new(service.Work)

	//get key word
	keyword := fmt.Sprintf("%v", json["keyword"])

	//query work list
	result := work.Search(keyword)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: err.Error()}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

//New comment
func Comment(c *gin.Context) {
	logger := utils.Log()
	logger.Info("Creat comment")
	var comment service.Comment
	//ID text
	err := c.BindJSON(&comment)
	logger.Info(err)
	logger.Info(comment)
	comment.IsNew = true
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	result := comment.Insert()
	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "Submit error"}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 0, Msg: "successffff"}
	res.Json(c)
}
