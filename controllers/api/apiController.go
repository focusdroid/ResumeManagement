package api

import (
	"ResumeManagement/helper"
	"ResumeManagement/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

type APIController struct{}

// SendMail
// @Tags 公共方法
// @Summary 邮箱获取验证码
// @Param mail query string true "mail"
// @Description do ping
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /sendMail [get]
func (api APIController) SendMail(c *gin.Context) {
	mail := c.Query("mail")
	status, _ := strconv.Atoi(c.DefaultQuery("status", "1"))
	var users = make(map[string]interface{})
	models.DB.Model(&models.User{}).Where("email = ?", mail).First(&users)
	if users["email"] == mail {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "该邮箱已被注册",
		})
		return
	}
	if !VerifyEmailFormat(mail) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"data":    gin.H{},
			"message": "邮箱格式异常",
		})
		return
	}
	SendCodeErr := helper.SendCode(mail, status)
	if SendCodeErr != nil {
		fmt.Println(SendCodeErr)
		return
	}
	code, err := helper.RedisGet(mail)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(code)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{
			"邮箱验证码": code,
		},
		"message": "发送成功验证码已发送" + mail + "邮箱,请及时查收",
	})
}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username query string true "username"
// @Param password query string true "password"
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /login [post]
func (api APIController) Login(c *gin.Context) {
	fmt.Println("login")
	json := make(map[string]interface{})

	c.ShouldBindJSON(&json)
	//username := strings.TrimSpace(json["username"].(string))
	//password := strings.TrimSpace(json["password"].(string))
	email := json["email"]
	password := json["password"]
	if email == nil || password == nil || email == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "用户名或密码不能为空",
			"data":    gin.H{},
		})
		return
	}
	//var user []*models.User
	//result := models.DB.Model(&models.User{}).Where("email = ?", email).First(&user)
	//fmt.Println(user[0].Email, user[0].Password)
	var users = make(map[string]interface{})
	err := models.DB.Model(&models.User{}).Where("email = ?", email).Find(&users).Error
	fmt.Println(reflect.TypeOf(users), reflect.TypeOf(users["email"]))
	if users["email"] == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "未查到该邮箱的注册记录",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "查询失败",
		})
		return
	}

	mail := helper.AutoRemoveSpace(users["email"].(string))
	passwords := helper.AutoRemoveSpace(users["pass_word"].(string))
	if mail != email || passwords != password {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"data":    gin.H{},
			"message": "用户名或密码错误",
		})
		return
	}
	// 返回指定字段
	type Userinfo struct {
		UserId string
		Name   string
		Email  string
		Phone  string
	}
	//var userlist *Userinfo
	var userlist = make(map[string]interface{})
	models.DB.Model(&models.User{}).Where("email = ?", email).Scan(&userlist)
	if userlist == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"data":    gin.H{},
			"message": "没有该用户信息，请核实账号重新登录",
		})
		return
	}
	fmt.Println(userlist["email"], userlist)
	token := helper.GenerateToken(&helper.UserClaims{
		UserID: userlist["user_id"].(string),
		Name:   userlist["name"].(string),
		Email:  userlist["email"].(string),
		Phone:  userlist["phone"].(string),
	})
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		//"data":    userlist,
		"token":   token,
		"message": "success",
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param username query string true "username"
// @Param password query string true "password"
// @Param code query string true "code"
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /register [post]
func (api APIController) Register(c *gin.Context) {
	json := make(map[string]interface{})
	//c.ShouldBindJSON(&json)
	c.BindJSON(&json)
	email := json["email"]
	password := json["password"]
	code := json["code"]

	fmt.Println(email, password, code)
	if email == nil || password == nil || code == nil || email == "" || password == "" || code == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "请将注册信息填写完整",
		})
		return
	}
	email = helper.AutoRemoveSpace(json["email"].(string))
	password = helper.AutoRemoveSpace(json["password"].(string))
	code = helper.AutoRemoveSpace(json["code"].(string))

	var userMessage = make(map[string]interface{})
	models.DB.Model(&models.User{}).Where("email = ?", email).First(&userMessage)
	if userMessage["email"] == email {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "该邮箱已经注册,请直接登录",
		})
		return
	}
	emailCode, err := helper.RedisGet(email.(string))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "验证码过期",
		})
		return
	}
	var user []*models.User
	dataEmail := models.DB.Model(&models.User{}).First(&user, "email = ?", email)
	fmt.Println(dataEmail.RowsAffected, dataEmail)
	if dataEmail.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "该邮箱已经被注册",
		})
		return
	}
	fmt.Println(emailCode, code, emailCode != code)
	if emailCode != code {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "验证码输入错误,请重新获取验证码",
		})
		helper.RedisDelete(email.(string))
		return
	}
	users := models.User{
		Email:    email.(string),
		UserId:   email.(string),
		Password: password.(string),
	}
	createErr := models.DB.Model(&models.User{}).Create(&users).Error
	if createErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"data":    gin.H{},
			"message": "注册失败，请重试",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"data":    gin.H{},
		"message": "注册成功",
	})
	val, err := helper.RedisDelete(email.(string)) // 注册成功之后删除redis中的随机码
	if err != nil {
		fmt.Println(val, err)
		helper.RedisDelete(email.(string))
	}
}

// @Tags 公共方法
// @Summary 刷新token
// @Param token query string true "token"
// @Description do ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":"200", "message":"", "data":""}"
// @Router /refreshToken [get]
func (api APIController) RefreshToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "参数不能为空",
		})
		return
	}
	//fmt.Println(token)
	claims, err := helper.ParseToken(c, token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    "-1",
			"message": "token失效",
		})
		return
	}
	//fmt.Println("===", claims.ExpiresAt)
	currentTime := time.Now().Unix()
	if claims.ExpiresAt < currentTime {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2001",
			"message": "token失效，请重新登录",
		})
		return
	}
	jwt.TimeFunc = time.Now
	claims.ExpiresAt = time.Now().Add(helper.EffectTime).Unix()
	newToken := helper.GenerateToken(claims)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"token":   newToken,
	})
}
