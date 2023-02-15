package backlog

import (
	"ResumeManagement/helper"
	"ResumeManagement/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type BacklogController struct{}

// AddBacklog
// @Tags 用户相关方法
// @Summary 新增待办信息
// @Param backlogText query string true "backlogText"
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /backlog/addBacklog [post]
func (back BacklogController) AddBacklog(c *gin.Context) {
	json := make(map[string]interface{})
	c.ShouldBindJSON(&json)
	if json["backlogText"] == nil || json["backlogText"] == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "请将信息填写完整",
		})
		return
	}
	user, _ := helper.AnalysisTokenGetUserInfo(c)
	var users = make(map[string]interface{})
	err := models.DB.Model(&models.User{}).Where("email = ?", user.Email).First(&users).Error
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "查询异常",
		})
		return
	}
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

	backlogText := json["backlogText"].(string)

	backlog := models.Backlog{
		UserId:      users["email"].(string),
		BacklogText: backlogText,
	}

	createErr := models.DB.Model(models.Backlog{}).Create(&backlog).Error
	if createErr != nil {
		fmt.Println("createErr", createErr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "插入数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "待办信息新增成功！",
	})
}

// GetBacklogList
// @Tags 用户相关方法
// @Summary 获取待办信息列表
// @Description { backlog_type: 1/2 } 1 正在待办  2 已经完成(有效期1个月内的，按照创建时间获取)
// @Param backlog_type query string true "backlog_type"
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /backlog/getBacklogList [get]
func (back BacklogController) GetBacklogList(c *gin.Context) {
	user, _ := helper.AnalysisTokenGetUserInfo(c)
	fmt.Println(user.Email)
	bl_type := c.Query("backlog_type")
	if bl_type == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "backlog_type不能为空",
			"data":    gin.H{},
		})
		return
	}
	backlog_type, typeError := strconv.Atoi(c.Query("backlog_type"))
	if typeError != nil {
		fmt.Println("typeError", typeError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "查询数据库异常",
			"data":    gin.H{},
		})
		return
	}
	var order_text string
	if backlog_type == 1 { // 创建的使用创建时间排序
		order_text = "created_at desc"
	} else if backlog_type == 2 { // 已完成的使用最后的更新时间排序
		order_text = "updated desc"
	}
	var backlogList []models.BacklogInterface
	err := models.DB.Model(models.Backlog{}).Where("user_id = ? and backlog_status != ? and backlog_type = ?", user.Email, 0, backlog_type).Order(order_text).Find(&backlogList).Error
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
		"code":         "200",
		"message":      "success",
		"data":         backlogList,
		"backlog_type": backlog_type,
	})
}

// ChangeBackStatus
// @Tags 用户相关方法
// @Summary 删除/置为已完成/置为未完成
// @Description { BacklogStatus: 0/1/2/3/4 } 已删除 0 正常显示1 轻度紧急2 中度紧急3 非常紧急4
// @Description { BacklogType: 1/2 } 正在待办1,已完成2
// @Param backlog_type query string true "backlog_type"
// @Param backlog_status query string true "backlog_type"
// @Param id query string true "id"
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /backlog/getBacklogList [get]
func (back BacklogController) ChangeBackStatus(c *gin.Context) {
	json := make(map[string]int)
	c.ShouldBindJSON(&json)

	id := json["id"]
	backlog_status := json["backlog_status"]
	backlog_type := json["backlog_type"]
	fmt.Println(id, backlog_status, backlog_type)
	fmt.Println(reflect.TypeOf(id), reflect.TypeOf(backlog_status), reflect.TypeOf(backlog_type))
	backlog_status_list := []int{0, 1, 2, 3, 4}
	backlog_type_list := []int{1, 2}

	for key, _ := range backlog_type_list {
		if key != backlog_type {
			c.JSON(http.StatusOK, gin.H{
				"code":    "-1",
				"message": "该参数不符合要求",
			})
			return
		}
	}
	for key, _ := range backlog_status_list {
		if key != backlog_status {
			c.JSON(http.StatusOK, gin.H{
				"code":    "-1",
				"message": "该参数不符合要求",
			})
			return
		}
	}
	models.DB.Model(&models.Backlog{})

}
