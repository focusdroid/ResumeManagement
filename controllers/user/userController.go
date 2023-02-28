package user

import (
	"ResumeManagement/helper"
	"ResumeManagement/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct{}

// UserList
// @Tags 管理员(admin)方法
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

	var userList []models.UserField
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
// @Tags 管理员(admin)方法
// @Summary 管理员手动添加用户
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /user/addUser [get]
func (user UserController) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    "",
	})
}

// UserInfo
// @Tags 用户相关方法
// @Summary 获取用户信息
// @Description /user/userinfo
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /user/userinfo [post]
func (user UserController) UserInfo(c *gin.Context) {
	users, userError := helper.AnalysisTokenGetUserInfo(c)
	if userError != nil {
		fmt.Println("userError", userError)
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "用户token异常",
		})
		return
	}
	email := users.Email
	var userinfo models.UserField
	userFindErr := models.DB.Model(&models.User{}).Where("is_delete = ?", 0).Where("email = ?", email).First(&userinfo).Error
	if userFindErr != nil {
		fmt.Println("userFindErr", userFindErr)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "userFindErr",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    userinfo,
	})
	return
}
