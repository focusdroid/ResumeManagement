package routers

import (
	"ResumeManagement/controllers/list"
	"ResumeManagement/middleware"
	"github.com/gin-gonic/gin"
)

func ListInfoRouters(r *gin.Engine) {
	listInfoRouters := r.Group("/list", middleware.InitMiddleware, middleware.InitMiddlewareBlacklist)
	{
		listInfoRouters.GET("/resume", list.ListController{}.ResumeList)            // 获取简历列表
		listInfoRouters.GET("/mainResume", list.ListController{}.MainResumeList)    // 获取重点关注人群简历
		listInfoRouters.POST("/addUserResume", list.ListController{}.AddUserResume) // 添加用户信息(简历)
		listInfoRouters.POST("/upload", list.ListController{}.Upload)               // 上传文件[暂时废弃]
		listInfoRouters.POST("/modifyMain", list.ListController{}.ModifyMainStatus) // 取消/添加重点标记
		listInfoRouters.GET("/detail", list.ListController{}.ResumeDetail)          // 获取简历详情
		listInfoRouters.GET("/delete", list.ListController{}.ResumeDelete)          // 删除个人信息
		listInfoRouters.GET("/deleted", list.ListController{}.ResumeDeleted)        // 获取已经删除简历信息(管理员权限)
		listInfoRouters.POST("/updateInfo", list.ListController{}.UpdateResumeInfo) // 更新简历信息
	}
}
