package helper

import (
	"ResumeManagement/models"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/http"
	"net/smtp"
	"regexp"
	"strconv"
	"time"
)

// 生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// 邮箱发送验证码
func SendCode(toUserEmail string, code int) error {
	e := email.NewEmail()
	e.From = "Get <focusdroid@163.com>"
	e.To = []string{toUserEmail}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "验证码发送测试"
	var num string = GetRand()
	fmt.Println(num)
	var text string
	if code == 0 {
		text = "注册"
	} else if code == 1 {
		text = "登录"
	}
	//e.Text = []byte("您的验证码是：<b>123456</b>")
	e.HTML = []byte("系统" + text + "的验证码是：<b>" + num + "</b>")
	//err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "focusdroid@163.com", "MYBVJUDOLMJTSERW", "smtp.163.com"))
	// 返回EOF时候，关闭SSL重试, 授权需要获取授权码
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "focusdroid@163.com", "MYBVJUDOLMJTSERW", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	RedisSet(toUserEmail, num)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// redis缓存
var ctx = context.Background()

func RedisSet(key string, value string) error {
	err := models.RDB.Set(ctx, key, value, time.Second*10*6*10)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

// 获取redis值
func RedisGet(key string) (string, error) {
	val, err := models.RDB.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return val, nil
}

// 删除redis
func RedisDelete(key string) (string, error) {
	val, err := models.RDB.Del(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(val), nil
}

// 去除空格
func AutoRemoveSpace(key string) string {
	if key == "" {
		return key
	}
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(key, "")
}

type UserClaims struct {
	UserID string
	Name   string
	Email  string
	Phone  string
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f487")
	//该路由下不校验token
	//noVerify = []interface{}{"/login", "/ping"}
	noVerify = []string{"/login", "/register", "/sendMail", "/swagger/index.html"}
	//token有效时间（纳秒）
	EffectTime = 2 * time.Hour
)

// 生成token
func GenerateToken(claims *UserClaims) string {
	claims.ExpiresAt = time.Now().Add(EffectTime).Unix()
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign
}

// 验证token
func JwtVerify(url string) bool {
	//过滤是否验证token
	//文档里我没给出utils.IsContainArr的代码，这个东西不重要，你直接删掉这段都行，这只是一个url过滤的逻辑
	/*if utils.IsContainArr(noVerify, c.Request.RequestURI) {
		return
	}*/
	if isHave(noVerify, url) {
		return true
	}
	return false
	/*token := c.GetHeader("token")
	if token == "" {
		panic("token not exist !")
	}*/
	//验证token，并存储在请求中
	//c.Set("user", parseToken(token))
}

/*func JwtVerify(c *gin.Context) {
	//过滤是否验证token
	//文档里我没给出utils.IsContainArr的代码，这个东西不重要，你直接删掉这段都行，这只是一个url过滤的逻辑
	/*if utils.IsContainArr(noVerify, c.Request.RequestURI) {
		return
	}
	if isHave(noVerify, c.Request.RequestURI) {
		return
	}
	token := c.GetHeader("token")
	if token == "" {
		panic("token not exist !")
	}
	//验证token，并存储在请求中
	//c.Set("user", parseToken(token))
}*/

// 解析Token
/*func parseToken(tokenString string) *UserClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}
	return claims
}*/

// 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}

func ParseToken(c *gin.Context, tokenString string) (*UserClaims, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	//fmt.Println(token)
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return &UserClaims{}, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2001",
			"message": "token解析出错2",
		})
		c.Abort()
		return &UserClaims{}, err
		//panic("token is valid")
	}
	return claims, nil

	//jwt.TimeFunc = time.Now
	//claims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	//return GenerateToken(claims)
}

/*
	func Refresh(tokenString string) string {
		jwt.TimeFunc = func() time.Time {
			return time.Unix(0, 0)
		}
		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			panic(err)
		}
		claims, ok := token.Claims.(*UserClaims)
		if !ok {
			panic("token is valid")
		}
		jwt.TimeFunc = time.Now
		claims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
		return GenerateToken(claims)
	}
*/
func isHave(noVerify []string, url string) bool {
	for _, v := range noVerify {
		if v == url {
			return true
		}
	}
	return false
}

// 对token进行解析
func AnalysisTokenGetUserInfo(c *gin.Context) (*UserClaims, error) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    "2001",
			"message": "未携带token",
		})
	}
	userinfo, err := ParseToken(c, token)
	return userinfo, err
}

// 异常返回封装
func AbnormalEncapsulation(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    "-1",
		"data":    gin.H{},
		"message": message,
	})
}
