package api

import (
	"po_go/service"
	"po_go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Find the comment list for a work
func ChatRoomList(c *gin.Context) {
	logger := utils.Log()
	logger.Info("Get ChatRoom List")
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	logger.Info(json)
	chatRoomList := new(service.ChatRoomAlias)

	//get member Id
	memberID, err := strconv.Atoi(utils.StrVal(json["member_id"]))
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Query fail", Data: ""}
		res.Json(c)
	}
	logger.Info(memberID)

	//query comment list
	result := chatRoomList.FindChatRoomAliasByMemberID(memberID)
	logger.Info(result)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: err.Error()}
		res.Json(c)
		return
	}

	res := &utils.Response{Code: 200, Msg: "", Data: result}
	res.Json(c)
}
