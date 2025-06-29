package middleware

import (
	"bytes"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestMiddleware struct{}

func NewRequestMiddleware() *RequestMiddleware {
	return &RequestMiddleware{}
}

func (*RequestMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 读取请求体（并还原给后续读取）
		var bodyBytes []byte
		if r.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(r.Body)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 记录请求信息
		logx.Infof("[Request] %s %s - Body: %s", r.Method, r.URL.Path, string(bodyBytes))

		// 响应捕获
		rw := &responseWriter{ResponseWriter: w, body: &bytes.Buffer{}}

		// 调用下一个 handler
		next(rw, r)

		// 日志记录响应和耗时
		duration := time.Since(start)
		logx.Infof("[Response] %s %s - Code: %d - Body: %s - Time: %v",
			r.Method, r.URL.Path, rw.code, rw.body.String(), duration)
	}
}

type responseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
	code int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b) // 缓存响应
	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.code = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
