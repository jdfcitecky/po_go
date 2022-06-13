package middleware

import (
	"io/ioutil"
	"po_go/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {

	logger := utils.Log()

	return func(c *gin.Context) {
		// startTime
		startTime := time.Now()

		// go to next middleware
		c.Next()

		// endtime
		endTime := time.Now()

		// latencyTime
		latencyTime := endTime.Sub(startTime) / time.Millisecond

		// reqMethod
		reqMethod := c.Request.Method

		// reqUri
		reqUri := c.Request.RequestURI

		header := c.Request.Header
		proto := c.Request.Proto

		// statusCode
		statusCode := c.Writer.Status()

		// clientIP
		clientIP := c.ClientIP()

		err := c.Err()

		body, _ := ioutil.ReadAll(c.Request.Body)

		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"proto":        proto,
			"err":          err,
			"body":         body,
		}).Info()
	}
}
