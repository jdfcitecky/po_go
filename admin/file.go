package admin

import (
	"fmt"
	"path"
	"po_go/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//find commentt list to manage panel
func Upload(c *gin.Context) {
	logger := utils.Log()
	logger.Info("File upload")
	// get file from form input name 'file'
	file, err := c.FormFile("file")
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "file upload error"}
		res.Json(c)
		return
	}
	// createpath
	now := time.Now()
	ext := path.Ext(file.Filename)
	fileName := strconv.Itoa(now.Nanosecond()) + ext
	filePath := fmt.Sprintf("%s%s%s%s",
		"./static/upload/",
		fmt.Sprintf("%04d", now.Year()),
		fmt.Sprintf("%02d", now.Month()),
		fmt.Sprintf("%04d", now.Day()))
	utils.CreateDir(filePath)
	fullPath := filePath + "/" + fileName

	// save file to tmp folder in current directory
	c.SaveUploadedFile(file, fullPath)
	url := (fullPath[1:])
	res := &utils.Response{Code: 0, Msg: "", Data: url}
	res.Json(c)

}
