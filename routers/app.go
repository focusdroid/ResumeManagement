package routers

import (
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	r := gin.New()
	r.Use(middleware.Loggoer())
	r.Use(gin.Recovery())
	DocsRouter(r)      // 文档地址
	ApiRouters(r)      // 公共的api
	ListInfoRouters(r) // 简历列表接口
	UserRouter(r)      // 用户信息
	BacklogRouters(r)  // 待办信息

	return r
}
