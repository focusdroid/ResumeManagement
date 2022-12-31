package user

import (
	"ResumeManagement/helper"
	"ResumeManagement/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserController struct{}

// UserList
// @Tags admin方法
// @Summary 获取系统内所有的用户
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /user/userList [get]
func (user UserController) UserList(c *gin.Context) {
	page, pageError := strconv.Atoi(c.DefaultQuery("page", "1"))
	if pageError != nil {
		fmt.Println("pageError", pageError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "page参数异常",
			"data":    "",
		})
	}
	pageSize, pageSizeError := strconv.Atoi(c.DefaultQuery("pageSize", "30"))
	if pageSizeError != nil {
		fmt.Println("pageError", pageError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "pageSize参数异常",
			"data":    "",
		})
	}
	userinfo, err := helper.AnalysisTokenGetUserInfo(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userinfo)
	var userAdmin models.User
	findError := models.DB.Model(models.User{}).Where("user_id = ?", userinfo.UserID).First(&userAdmin).Error
	if findError != nil {
		fmt.Println(findError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "查询数据异常",
			"data":    gin.H{},
		})
		return
	}
	if !userAdmin.IsAdmin {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "当前用户不是管理员",
			"data":    gin.H{},
		})
		return
	}
	pageNumber := (page - 1) * pageSize
	//var userList []models.User
	type limitField struct {
		ID         uint       `json:"id"`
		UserId     string     `json:"user_id"`
		Name       string     `json:"name"`
		Phone      string     `json:"phone"`
		Email      string     `json:"email"`
		Gender     uint       `json:"gender"`
		AvatarUrl  string     `json:"avatar_url"`
		Country    string     `json:"country"`
		Province   string     `json:"province"`
		City       string     `json:"city"`
		NickName   string     `json:"nick_name"`
		UserStatus string     `json:"user_status"`
		IsAdmin    bool       `json:"is_admin"`
		LineTime   *time.Time `json:"line_time"`
	}

	var userList []limitField
	FindError := models.DB.Model(&models.User{}).Offset(pageNumber).Limit(pageSize).Where("is_admin = ?", false).Scan(&userList).Error
	if FindError != nil {
		fmt.Println(FindError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "查询数据异常111",
			"data":    gin.H{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data": gin.H{
			"data":        userList,
			"total":       len(userList),
			"currentPage": page,
			"pageSize":    pageSize,
		},
	})
}

// AddUser
// @Tags admin方法
// @Summary 管理员手动添加用户
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /user/userList [get]
func (user UserController) AddUser(c *gin.Context) {

}
