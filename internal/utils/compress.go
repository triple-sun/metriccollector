package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer

	gz := gzip.NewWriter(&b)

	defer gz.Close()

	if _, err := gz.Write(data); err != nil {
		return nil, fmt.Errorf("error compressing data: %v", err)
	}

	if err := gz.Close(); err != nil {
		return nil, fmt.Errorf("failed compress data: %v", err)
	}
	return b.Bytes(), nil
}

func ShouldServerCompress(c *gin.Context) bool {
	return (strings.Contains(c.GetHeader("Content-Type"), "application/json") || strings.Contains(c.GetHeader("Content-Type"), "text/html")) && strings.Contains(c.GetHeader("Accept-Encoding"), "gzip")
}
