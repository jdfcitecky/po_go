package router

import (
	"po_go/admin"
	"po_go/api"

	"po_go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()

	//recovery from crashing
	router.Use(gin.Recovery())
	//to storage the image as file server
	router.Static("/static", "static")

	//cross middleware
	router.Use(middleware.Cors())

	//log middleware
	router.Use(middleware.Logger())

	register(router)

	return router

}

func register(router *gin.Engine) {
	//Find manager
	router.POST("/manager", api.FindManager)
	//Search
	router.POST("/search", api.Search)
	//Find a work
	router.POST("/work/show", api.FindWork)
	//comment
	router.POST("/work/comment", api.Comment)
	//comment
	router.POST("/work/comment/list", api.CommentListForWork)

	//member
	router.POST("/login", api.Login)
	router.POST("/logout", api.Logout)
	router.POST("/member/update", api.MemberUpdateInfo)
	// chat room
	router.GET("/ws/:chatroomid/:senderid", api.ChatRoomSocketHandler)
	idVerify := router.Group("/chatroom", middleware.IdVerify())
	{
		idVerify.POST("/list", api.ChatRoomList)
		idVerify.POST("/message/list", api.ChatRoomMessagesList)
		idVerify.POST("/message/list/all", api.ChatRoomMessagesList)
		idVerify.POST("/message/save", api.ChatRoomMessagesSave)
		idVerify.POST("/message/update/read", api.ChatRoomMessagesUpdateRead)
		// idVerify.POST("/ws/:id", api.ChatRoomSocketHandler)
	}
	//admin
	jwt := router.Group("/admin", middleware.Jwt())
	{
		jwt.POST("/statistic/get", admin.GetStatisticData)
		jwt.POST("/upload", admin.Upload)
		jwt.POST("/work/list", admin.WorkList)
		jwt.POST("/work/one", admin.FindWorkForManage)
		jwt.POST("/work/save", admin.WorkSave)
		jwt.POST("/work/delete", admin.WorkDelete)
		jwt.POST("/comment/list", admin.CommentList)
		jwt.POST("/comment/review", admin.CommentReview)
		jwt.POST("/comment/delete", admin.CommentDelete)
		// jwt.POST("/", admin.CommentDelete)
	}

}
