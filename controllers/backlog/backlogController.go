package backlog

import (
	"ResumeManagement/helper"
	"ResumeManagement/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BacklogController struct{}

// AddBacklog
// @Tags 用户相关方法
// @Summary 编辑待办信息
// @Param backlogText query string true "backlogText"
// @Param backlogStatus query string true "backlogStatus"
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /backlog/addBacklog [post]
func (back BacklogController) AddBacklog(c *gin.Context) {
	token := c.GetHeader("token")
	user, _ := helper.ParseToken(c, token)
	var users = make(map[string]interface{})
	models.DB.Model(&models.User{}).Where("email = ?", user.Email).First(&users)
	if users == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "获取用户信息失败",
		})
		return
	}
	fmt.Println(users["email"], user.Email)
	if users["email"] != user.Email {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "获取用户信息失败",
		})
		return
	}

	json := make(map[string]interface{})
	c.ShouldBindJSON(&json)
	if json["backlogText"] == nil || json["backlogText"] == "" || json["backlogStatus"] == nil || json["backlogStatus"] == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "请将信息填写完整",
		})
	}
	backlogText := json["backlogText"].(string)
	backlogStatus := json["backlogStatus"].(int)

	backlog := models.Backlog{
		UserId:        users["email"].(string),
		BacklogText:   backlogText,
		BacklogStatus: backlogStatus,
	}

	err := models.DB.Model(models.Backlog{}).Create(&backlog).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "插入数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "添加信息成功！",
	})
}

// GetBacklogList
// @Tags 用户相关方法
// @Summary 获取待办信息列表
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /backlog/getBacklogList [get]
func (back BacklogController) GetBacklogList(c *gin.Context) {
	token := c.GetHeader("token")
	user, _ := helper.ParseToken(c, token)
	fmt.Println(user.Email)

	var backlogList []models.BacklogInterface
	err := models.DB.Model(models.Backlog{}).Where("user_id = ? && backlog_status != ?", user.Email, 0).Find(&backlogList).Error
	//err := models.DB.Model(models.Backlog{}).Where("user_id = ?", user.Email).Find(&backlogList).Error
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "查询数据库异常",
			"data":    gin.H{},
		})
		return
	}
	fmt.Println(backlogList)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    backlogList,
	})
}
