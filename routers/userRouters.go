package routers

import (
	"ResumeManagement/controllers/user"
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

/**
 * @name: focusdroid
 * @description: 用户相关
 * @author: 1.0
 * @time：2023-02-28 19:41:00
**/

func UserRouter(r *gin.Engine) {
	userRouter := r.Group("user", middleware.InitMiddleware, middleware.InitMiddlewareBlacklist)
	{
		userRouter.GET("/userList", user.UserController{}.UserList) // 获取所有系统用户(管理员和主管可以调用)
		userRouter.GET("/addUser", user.UserController{}.AddUser)   // 主管可以手动添加用户(管理员和主管可以调用)
		userRouter.GET("/userinfo", user.UserController{}.UserInfo) // 获取用户信息接口
	}
}
