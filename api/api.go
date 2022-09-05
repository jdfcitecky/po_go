package api

import (
	"fmt"
	"po_go/service"
	"po_go/utils"
	"strconv"
	"strings"

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
	logger := utils.Log()
	logger.Info("Get Comments")
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	logger.Info(json)
	comment := new(service.Comment)

	//get work Id
	workID, err := strconv.Atoi(utils.StrVal(json["work_id"]))
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Query fail", Data: ""}
		res.Json(c)
	}
	pageStart, err := strconv.Atoi(utils.StrVal(json["page_start"]))
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Query fail", Data: ""}
		res.Json(c)
	}
	pageLimit, err := strconv.Atoi(utils.StrVal(json["page_limit"]))
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Query fail", Data: ""}
		res.Json(c)
	}
	//query comment list
	result := comment.FindCommentsListByWorkID(workID, pageStart, pageLimit)
	count := comment.FindNumberOfCommentsByWorkID(workID)
	logger.Info(result)
	logger.Info(count)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: err.Error()}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 200, Msg: "", Data: result, Count: count}
	res.Json(c)
}

//Find a work
func FindWork(c *gin.Context) {
	logger := utils.Log()
	logger.Info("Find work")
	var work service.Work
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
	// comments := comment.FindCommentsListByWorkID(work.ID)
	Map := make(map[string]interface{})
	Map["work"] = result
	// Map["comments"] = comments
	res := &utils.Response{Code: 0, Msg: "", Data: Map}

	//write a record
	ip := c.ClientIP()
	workID := work.ID
	browseRecord := service.BrowseRecord{
		WorkID: workID,
		IP:     ip,
		Date:   utils.GetCurrentDate(),
	}
	browseRecord.Insert()

	res.Json(c)
}

//Find the work list
func Search(c *gin.Context) {
	logger := utils.Log()
	logger.Info("Search")
	json := make(map[string]interface{})
	err := c.ShouldBind(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	work := new(service.Work)

	//get key word
	keywords := fmt.Sprintf("%s", json["keyWord"])
	words := strings.Fields(keywords)
	result := make([]*service.Work, 0)
	for _, keyword := range words {
		//query work list
		resultInOne := work.Search(keyword)
		result = append(result, resultInOne...)

	}
	// jsut a wrapper
	Map := make(map[string]interface{})
	Map["works"] = result
	// result["works"] = resultInOne
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

	res := &utils.Response{Code: 0, Msg: "success"}

	//write a record
	workID := comment.WorkID
	memberID := comment.MemberID
	ip := c.ClientIP()
	commentRecord := service.CommentRecord{
		MemberID: memberID,
		WorkID:   workID,
		IP:       ip,
		Date:     utils.GetCurrentDate(),
	}
	commentRecord.Insert()
	res.Json(c)
}
