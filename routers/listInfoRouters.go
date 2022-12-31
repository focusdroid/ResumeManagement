package routers

import (
	"ResumeManagement/controllers/list"
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

func ListInfoRouters(r *gin.Engine) {
	listInfoRouters := r.Group("/list", middleware.InitMiddleware)
	{
		listInfoRouters.GET("/resume", list.ListController{}.ResumeList)            // 获取简历列表
		listInfoRouters.GET("/mainResume", list.ListController{}.MainResumeList)    // 获取重点关注人群简历
		listInfoRouters.POST("/addUserResume", list.ListController{}.AddUserResume) // 添加用户信息(简历)
		listInfoRouters.POST("/upload", list.ListController{}.Upload)               // 上传文件[暂时废弃]
		listInfoRouters.POST("/modifyMain", list.ListController{}.ModifyMainStatus) // 取消/添加重点标记
	}
}
