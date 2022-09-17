package middleware

import (
	"po_go/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func IdVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := utils.Log()
		logger.Info("--------------------------------in jwt Verify")

		token := c.GetHeader("token")
		if token == "" {
			logger.Info("--------------------------------no token Verify")
			res := &utils.Response{
				Code: 1100,
				Msg:  "Request without token",
			}
			res.Json(c)
			c.Abort()
			return
		}
		logger.Info("--------------------------------in idverify ", token)
		logrus.Debug("get Token:", token)

		data, found := utils.Cache.Get(token)
		if found == false {
			res := &utils.Response{
				Code: 1100,
				Msg:  "token can't find or expired",
			}
			res.Json(c)
			c.Abort()
			return
		}

		tokenData := data.(*TokeData)

		b := token != tokenData.SetToken()
		if b {
			res := &utils.Response{
				Code: 1100,
				Msg:  "token error",
			}
			res.Json(c)
			c.Abort()
			return
		}

		c.Set("token", data)
		//claims = c.MustGet("claims").(*CustomClaims)

		c.Next()
	}
}
