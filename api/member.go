package api

import (
	"po_go/cache"
	"po_go/middleware"
	"po_go/service"
	"po_go/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//sign in
func Login(c *gin.Context) {
	logger := utils.Log()
	logger.Info("Login")
	var member service.Member
	err := c.BindJSON(&member)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}

	result := member.Login()

	if result.ID == 0 {
		res := &utils.Response{Code: 1001, Msg: "Can not find member", Data: "Can not find member"}
		res.Json(c)
		return
	}

	if result.Password != utils.Md5(member.Password) {
		res := &utils.Response{Code: 1002, Msg: "Member password wrong", Data: "Member password wrong"}
		res.Json(c)
		return
	}

	key := strconv.Itoa(time.Now().Nanosecond())

	token := &middleware.TokeData{
		Username: utils.Md5(member.Email),
		Password: utils.Md5(member.Password),
		Key:      key,
	}

	tokenKey := token.SetToken()
	Map := make(map[string]interface{})
	result.Password = ""
	// Map["member"] = result
	Map["tokenKey"] = tokenKey
	Map["email"] = result.Email
	Map["memberID"] = result.ID
	if result.IsManager == true {
		Map["isManager"] = true
	}
	//put token into cache
	utils.Cache.Set(tokenKey, token, cache.DefaultExpiration)
	res := &utils.Response{Code: 0, Msg: "",
		Data: Map}
	res.Json(c)
}

//sign out
func Logout(c *gin.Context) {
	token := c.GetHeader("token")
	utils.Cache.Delete(token)
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)
}

//modify member info
func MemberUpdateInfo(c *gin.Context) {
	logger := utils.Log()

	var member service.Member
	err := c.BindJSON(&member)
	logger.Info(member)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	if member.Password != "" {
		member.Password = utils.Md5(member.Password)
	}
	var result *gorm.DB
	if member.ID <= 0 {
		var chatRoom service.ChatRoom
		var chatRoomAlias service.ChatRoomAlias
		var chatRoomMember service.ChatRoomMember
		// Create a new chat room and the chat room id is same to member id
		result = chatRoom.Insert()
		if result.Error != nil {
			res := &utils.Response{Code: 1000, Msg: "Submit error"}
			res.Json(c)
			return
		}
		newID := member.FindNewMemberID() + 1
		// For this member can send to admin
		result = chatRoomAlias.Insert(newID, newID, "Admin - Hsin Cho")
		if result.Error != nil {
			res := &utils.Response{Code: 1000, Msg: "Submit error"}
			res.Json(c)
			return
		}

		// For this admin can send to this member
		result = chatRoomAlias.Insert(1, newID, member.Email)
		if result.Error != nil {
			res := &utils.Response{Code: 1000, Msg: "Submit error"}
			res.Json(c)
			return
		}

		// Put this member to the new chat room
		result = chatRoomMember.Insert(newID, newID)
		if result.Error != nil {
			res := &utils.Response{Code: 1000, Msg: "Submit error"}
			res.Json(c)
			return
		}

		// Put admin to the new chat room
		result = chatRoomMember.Insert(1, newID)
		if result.Error != nil {
			res := &utils.Response{Code: 1000, Msg: "Submit error"}
			res.Json(c)
			return
		}

		// create new member
		result = member.Insert()
		if result.Error != nil {
			res := &utils.Response{Code: 1000, Msg: "Submit error"}
			res.Json(c)
			return
		}
	} else {
		result = member.UpdateInfo()
	}

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "Submit error"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)

}

//modify password
func BloggerUpdatePassword(c *gin.Context) {
	var member service.Member
	err := c.BindJSON(&member)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "Data format wrong"}
		res.Json(c)
		return
	}
	if member.Password != "" {
		member.Password = utils.Md5(member.Password)
	}
	var result *gorm.DB
	if member.ID <= 0 {
		result = member.Insert()
	} else {
		result = member.UpdatePassword()
	}

	if result.Error != nil {
		res := &utils.Response{Code: 1000, Msg: "Submit error"}
		res.Json(c)
		return
	}
	res := &utils.Response{Code: 0, Msg: ""}
	res.Json(c)

}
