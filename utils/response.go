package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//define the struct of response
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Count int         `json:"count,omitempty"`
}

//return JSON data
func (res *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, res)
	return
}
