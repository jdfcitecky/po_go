package router

import (
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
	// //api
	// //查找博主
	// router.POST("/blogger", api.FindBlogger)
	// //查找博客类型
	// router.POST("/blog/type", api.FindType)
	// //查询博客列表
	// router.POST("/blog/list", api.BlogList)
	// //查询博客内容
	// router.POST("/blog/show", api.FindBlog)
	// //提交评论
	// router.POST("/blog/comment", api.Comment)

	// //admin
	// router.POST("/login", admin.Login)
	// router.POST("/logout", admin.Logout)
	// router.POST("/upload", admin.Upload)
	// router.POST("/uploadEdit", admin.UploadEdit)
	// jwt := router.Group("/admin", middleware.Jwt())
	// {
	// 	jwt.POST("/blogger/find", admin.FindBlogger)
	// 	jwt.POST("/blogger/updatePassword", admin.BloggerUpdatePassword)
	// 	jwt.POST("/blogger/updateInfo", admin.BloggerUpdateInfo)
	// 	jwt.POST("/blog/type/list", admin.TypeList)
	// 	jwt.POST("/blog/type/save", admin.BlogTypeSave)
	// 	jwt.POST("/blog/type/one", admin.TypeOne)
	// 	jwt.POST("/blog/type/delete", admin.BlogTypeDel)
	// 	jwt.POST("/blog/type/All", admin.TypeAll)
	// 	jwt.POST("/blog/save", admin.BlogSave)
	// 	jwt.POST("/blog/list", admin.BlogList)
	// 	jwt.POST("/blog/one", admin.BlogOne)
	// 	jwt.POST("/blog/delete", admin.BlogDelete)
	// 	jwt.POST("/comment/list", admin.CommentList)
	// 	jwt.POST("/comment/review", admin.CommentReview)
	// 	jwt.POST("/comment/delete", admin.CommentDelete)
	// }

}
