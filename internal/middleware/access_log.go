package middleware

import (
	"bytes"
	"time"

	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)

}

// 中间件对应的处理方法
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		//fmt.Println(bodyWriter)
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()

		// fmt.Println(c.Request.PostForm.Encode()) 输出 app_key=eddycjy&app_secret=go-programming-tour-book
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		//	fmt.Println(fields)
		//
		// global.Logger 是空~  fmt.Println(global.Logger)
		if global.Logger != nil {
			s := "access log: method: %s, status_code: %d, " +
				"begin_time: %d, end_time: %d"
				// 注意顺序, global.Logger 可能在这里用不了
			global.Logger.WithFields(fields).Infof(s,
				c.Request.Method,
				bodyWriter.Status(),
				beginTime,
				endTime,
			)
		}

	}
}
