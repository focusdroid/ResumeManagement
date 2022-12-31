package middleware

import (
	"ResumeManagement/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitMiddleware(c *gin.Context) {
	url := c.Request.RequestURI
	isNExt := helper.JwtVerify(url)
	if isNExt {
		return
	}
	token := c.GetHeader("token")
	fmt.Println(token, token == "")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "未携带token",
		})
		c.Abort()
		return
	}
	claims, err := helper.ParseToken(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2001",
			"message": "token解析出错3",
		})
		c.Abort()
		return
	}
	currentTime := time.Now().Unix()
	fmt.Println(currentTime, claims.ExpiresAt)
	if currentTime > claims.ExpiresAt {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2001",
			"message": "token失效，请重新登录,过滤器",
		})
		c.Abort()
		return
	}
}
