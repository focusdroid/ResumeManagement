package routers

import (
	"ResumeManagement/controllers/backlog"
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

func BacklogRouters(c *gin.Engine) {
	backlogRouters := c.Group("/backlog", middleware.InitMiddleware, middleware.InitMiddlewareBlacklist)
	{
		backlogRouters.POST("/addBacklog", backlog.BacklogController{}.AddBacklog)             // 添加待办
		backlogRouters.GET("/getBacklogList", backlog.BacklogController{}.GetBacklogList)      // 获取待办列表
		backlogRouters.POST("/changeBackStatus", backlog.BacklogController{}.ChangeBackStatus) // 修改待办状态
	}

}
