package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct{}

// GetUserList
// @Tags 用户相关方法
// @Summary 新增待办信息
// @Param backlogText query string true "backlogText"
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /backlog/addBacklog [post]
func (admin AdminController) GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    "",
	})
}
