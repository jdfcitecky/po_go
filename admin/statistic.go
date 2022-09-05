package admin

import (
	"po_go/service"
	"po_go/utils"

	"github.com/gin-gonic/gin"
)

//Find the comment list for a work
func GetStatisticData(c *gin.Context) {
	logger := utils.Log()
	logger.Info("Get statistic data")
	result := make(map[string]interface{})
	commentRecord := new(service.CommentRecord)
	browseRecord := new(service.BrowseRecord)
	// get datas
	dailyBrowse := browseRecord.FindNumberOfBrowseRecordGroupByDate()
	result["daily_browse"] = dailyBrowse
	dailyComment := commentRecord.FindNumberOfCommentRecordGroupByDate()
	result["daily_comment"] = dailyComment
	topBrowseWorks := browseRecord.FindNumberOfBrowseRecordGroupByWorkID()
	result["top_browse_works"] = topBrowseWorks
	topCommentWorks := commentRecord.FindNumberOfCommentRecordGroupByWorkID()
	result["top_comment_works"] = topCommentWorks

	logger.Info(result)

	res := &utils.Response{Code: 200, Msg: "", Data: result}
	res.Json(c)
}
