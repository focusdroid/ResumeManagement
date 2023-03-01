package middleware

import (
	"ResumeManagement/helper"
	"ResumeManagement/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	CurrentUserid string
	CurrentEmail  string
)

func InitMiddleware(c *gin.Context) {
	url := c.Request.RequestURI
	isNExt := helper.JwtVerify(url)
	fmt.Println("isNExt", isNExt, url)
	if isNExt {
		return
	}
	fmt.Println("isNExt2", isNExt, url)
	token := c.GetHeader("token")
	fmt.Println(token, token == "")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2001",
			"message": "未携带token",
		})
		c.Abort()
		return
	}
	claims, err := helper.ParseToken(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2001",
			"message": "token解析出错,请检查token是否有效",
		})
		c.Abort()
		return
	}
	userinfo, _ := helper.AnalysisTokenGetUserInfo(c)
	fmt.Println("userinfo", userinfo)
	CurrentUserid = userinfo.UserID
	CurrentEmail = userinfo.Email
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

func InitMiddlewareBlacklist(c *gin.Context) { // 黑名单
	url := c.Request.RequestURI
	isNExt := helper.JwtVerify(url)
	fmt.Println("InitMiddlewareBlacklistisNExt", isNExt, url)
	if isNExt {
		return
	}
	userinfo, _ := helper.AnalysisTokenGetUserInfo(c)
	fmt.Println("userinfo", userinfo)
	UserID := userinfo.UserID
	Email := userinfo.Email
	var user models.BlackList
	err := models.DB.Model(&models.BlackList{}).Where("email = ? or phone = ?", Email, UserID).Find(&user).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "2002",
			"message": err,
		})
		c.Abort()
		return
	}
	email := user.Email
	phone := user.Phone
	fmt.Println("------=====email !=phone ...", email != "", phone != "")
	if email != "" || phone != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2003",
			"message": "禁止该用户登录系统",
		})
		c.Abort()
		return
	}
}
