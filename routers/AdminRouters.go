package routers

import (
	"ResumeManagement/controllers/admin"
	"github.com/gin-gonic/gin"
)

func AdminRouters(c *gin.Engine) {
	adminRouters := c.Group("/admin")
	{
		adminRouters.GET("/getUserList", admin.AdminController{}.GetUserList)
	}
}
