package api

//find manager
// func FindManager(c *gin.Context) {
// 	var member service.Member
// 	result := member.Find()
// 	result.Password = ""
// 	res := &utils.Response{Code: 0, Msg: "", Data: result}
// 	res.Json(c)
// }

// //Find the comment list for a work
// func CommentList(c *gin.Context) {
// 	json := make(map[string]interface{})
// 	err := c.ShouldBind(&json)
// 	if err != nil {
// 		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
// 		res.Json(c)
// 		return
// 	}
// 	comment := new(service.Comment)
// 	//string to Int
// 	page, _ := strconv.Atoi(utils.StrVal(json["page"]))
// 	size, _ := strconv.Atoi(utils.StrVal(json["size"]))
// 	pageVo := &utils.Page{Page: page, Size: size, Total: blog.Count()}

// 	//get work Id
// 	workID, err := strconv.Atoi(utils.StrVal(json["work_id"]))
// 	if err == nil {
// 		comment.WorkID = workID
// 	}
// 	//查询博客列表
// 	result, err := comment.FindCommentListByWorkID(pageVo)
// 	if err != nil {
// 		res := &utils.Response{Code: 1000, Msg: err.Error()}
// 		res.Json(c)
// 		return
// 	}

// 	res := &utils.Response{Code: 0, Msg: "", Data: result, Count: pageVo.Total}
// 	res.Json(c)
// }

// //查询博客
// func FindBlog(c *gin.Context) {
// 	var blog service.Blog
// 	//绑定博客id值
// 	err := c.BindJSON(&blog)
// 	if err != nil {
// 		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
// 		res.Json(c)
// 		return
// 	}
// 	//更新点击次数
// 	blog.UpdateClick()

// 	//根据博客Id查询类型名
// 	result := blog.FindOneTypeName()

// 	//上一条
// 	last := blog.FindLastOne()
// 	//下一条
// 	next := blog.FindNextOne()

// 	//查找评论
// 	comments := blog.FindCommentByBlog()
// 	Map := make(map[string]interface{})
// 	Map["last"] = last
// 	Map["next"] = next
// 	Map["blog"] = result
// 	Map["comments"] = comments
// 	res := &utils.Response{Code: 0, Msg: "", Data: Map}
// 	res.Json(c)
// }

// //提交评论
// func Comment(c *gin.Context) {
// 	var comment service.Comment
// 	//ID text
// 	err := c.BindJSON(&comment)
// 	if err != nil {
// 		res := &utils.Response{Code: 1000, Msg: "数据格式出错"}
// 		res.Json(c)
// 		return
// 	}
// 	comment.Insert()

// 	res := &utils.Response{Code: 0, Msg: ""}
// 	res.Json(c)
// }
