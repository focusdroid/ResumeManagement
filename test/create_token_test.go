package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"testing"
	"time"
)

var mySigningKey = []byte("go-gin-resume")

type MyCustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"pass_word"`
	jwt.RegisteredClaims
}

// 生成token
func TestGenerateToken(t *testing.T) {
	userClaims := &MyCustomClaims{
		Name:             "tree",
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}

// 解析token
func TestAnalyseToken(t *testing.T) {
	toTokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidHJlZSJ9.te6hfNx6I8Y2v5iZoTUauh-OGVA6WKDikZK-4ZqJ914"
	userClaims := new(MyCustomClaims)
	claims, err := jwt.ParseWithClaims(toTokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(mySigningKey)

	if claims.Valid {
		fmt.Println(userClaims, userClaims.Name, userClaims.Password)
	}
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
	//noVerify = []string{"/login", "/ping"}
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
)

func TestRefresh(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJ3ZWV4c3NAMTYzLmNvbSIsIk5hbWUiOiIiLCJFbWFpbCI6IndlZXhzc0AxNjMuY29tIiwiUGhvbmUiOiIiLCJleHAiOjE2Njg4MzkwMjd9.HMviCj8rs-zPBPzW_EYg--_n27Ml39DKL3eNzMdE55Y"
	/*jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}*/

	token, _ := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	claims, ok := token.Claims.(*UserClaims)
	claims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	fmt.Println(claims, ok)
	fmt.Println(claims.Name, claims.Email, claims.Phone, claims.UserID, claims.ExpiresAt)
	//if err != nil {
	//	panic(err)
	//}
	//claims, ok := token.Claims.(*UserClaims)
	//if !ok {
	//	panic("token is valid")
	//}
	//jwt.TimeFunc = time.Now
	//claims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	//return GenerateToken(claims)
}
