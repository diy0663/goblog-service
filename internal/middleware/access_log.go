package middleware

import (
	"bytes"
	"time"

	"github.com/diy0663/goblog-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

		endTime := time.Now().Unix()

		logger.ZapLog.Info("access log:",
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.Int("status_code", bodyWriter.Status()),
			zap.Int("begin_time", int(beginTime)),
			zap.Int("end_time", int(endTime)),
			zap.String("request", c.Request.PostForm.Encode()),
			zap.String("response", bodyWriter.body.String()),
		)
	}
}
