package routers

import (
	"ResumeManagement/controllers/user"
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userRouter := r.Group("user", middleware.InitMiddleware)
	{
		userRouter.GET("/userList", user.UserController{}.UserList) // 获取所有系统用户(管理员和主管可以调用)
		userRouter.GET("/addUser", user.UserController{}.AddUser)   // 主管可以手动添加用户(管理员和主管可以调用)
	}
}
