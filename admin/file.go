package admin

import (
	"po_go/utils"

	"github.com/gin-gonic/gin"
)

//find commentt list to manage panel
func Upload(c *gin.Context) {
	logger := utils.Log()
	logger.Info("File upload")
	result := make(map[string]interface{})

	file, err := c.FormFile("file") // get file from form input name 'file'
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "file upload error"}
		res.Json(c)
		return
	}

	c.SaveUploadedFile(file, "./static/"+file.Filename) // save file to tmp folder in current directory
	result["url"] = ("static/" + file.Filename)
	res := &utils.Response{Code: 0, Msg: "", Data: result}
	res.Json(c)

	// file, err := c.FormFile("file")
	// if err != nil {
	// 	res := &utils.Response{Code: 1000, Msg: "文件上传错误"}
	// 	res.Json(c)
	// 	return
	// }

	// now := time.Now()
	// ext := path.Ext(file.Filename)
	// fileName := strconv.Itoa(now.Nanosecond()) + ext
	// filePath := fmt.Sprintf("%s%s%s%s",
	// 	"./static/upload/",
	// 	fmt.Sprintf("%04d", now.Year()),
	// 	fmt.Sprintf("%02d", now.Month()),
	// 	fmt.Sprintf("%04d", now.Day()))

	// utils.CreateDir(filePath)

	// fullPath := filePath + "/" + fileName
	// c.SaveUploadedFile(file, fullPath)
	// url := fullPath[1:len(fullPath)]
	// res := &utils.Response{Code: 0, Msg: "ok", Data: url}
	// res.Json(c)
	// return
}
