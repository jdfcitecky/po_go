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
	//admin
	jwt := router.Group("/admin", middleware.Jwt())
	{
		jwt.POST("/work/list", admin.WorkList)
		jwt.POST("/work/one", admin.FindWorkForManage)
		jwt.POST("/work/save", admin.WorkSave)
		jwt.POST("/blog/delete", admin.WorkDelete)
		jwt.POST("/comment/list", admin.CommentList)
		jwt.POST("/comment/review", admin.CommentReview)
		jwt.POST("/comment/delete", admin.CommentDelete)
	}

}
