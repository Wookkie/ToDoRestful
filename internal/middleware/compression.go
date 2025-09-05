package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GzipMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Content-Encoding") == "gzip" {
			gz, err := gzip.NewReader(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Не удалось распаковать gzip-запрос"})
				return
			}
			defer gz.Close()
			c.Request.Body = io.NopCloser(gz)
		}

		if strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
			c.Header("Content-Encoding", "gzip")

			gz := gzip.NewWriter(c.Writer)
			defer gz.Close()

			c.Writer = &gzipResponseWriter{
				ResponseWriter: c.Writer,
				Writer:         gz,
			}
		}

		c.Next()
	}
}

type gzipResponseWriter struct {
	gin.ResponseWriter
	io.Writer
}

func (w *gzipResponseWriter) Write(data []byte) (int, error) {
	contentType := w.Header().Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") || strings.HasPrefix(contentType, "text/html") {
		return w.Writer.Write(data)
	}
	return w.ResponseWriter.Write(data)
}
