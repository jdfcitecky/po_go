package admin

import (
	"po_go/service"
	"po_go/utils"

	"github.com/gin-gonic/gin"
)

//find commentt list to manage panel
func CommentList(c *gin.Context) {
	var comment service.Comment
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format error"}
		res.Json(c)
		return
	}
	result := comment.FindCommentListAll()
	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)
}

//Review comment
func CommentReview(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format error"}
		res.Json(c)
		return
	}
	result := comment.UpdateStatus()
	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "Review error"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}

//Delete comment
func CommentDelete(c *gin.Context) {
	var comment service.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format error"}
		res.Json(c)
		return
	}
	result := comment.DeleteComment()
	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "Delete error"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}
