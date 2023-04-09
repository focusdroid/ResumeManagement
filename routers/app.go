package routers

import (
	"ResumeManagement/middleware"
	"ResumeManagement/models"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	e := models.Enforcer

	e.LoadPolicy()
	/*r.POST("", func(c *gin.Context) {
		fmt.Println("增加Policy")
		if ok, _ := e.AddPolicy(e, "focusdroid_go@163.com", "admin", "/login", "GET"); !ok {
			fmt.Println("Policy已经存在")
		} else {
			fmt.Println("增加成功")
		}
	})*/
	r.Use(middleware.Loggoer())
	r.Use(gin.Recovery()) // Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500
	DocsRouter(r)         // 文档地址
	ApiRouters(r)         // 公共的api
	ListInfoRouters(r)    // 简历列表接口
	UserRouter(r)         // 用户信息
	BacklogRouters(r)     // 待办信息
	AdminRouters(r)       // 管理员
	return r
}
