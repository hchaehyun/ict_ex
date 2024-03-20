package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AllowedIp(allowIps string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		allowedIPsSlice := strings.Split(allowIps, ",") // 콤마를 기준으로 IP 주소를 구분합니다.
		for _, allowedIP := range allowedIPsSlice {
			allowedIP = strings.TrimSpace(allowedIP) // 공백 제거
			if clientIP == allowedIP {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		c.Abort()
	}
}
