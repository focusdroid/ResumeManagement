package test

import (
	"fmt"
	"testing"
)

var noVerify = []interface{}{"/login", "/ping"}
var url = "/login"

func TestJwtVerify(t *testing.T) {
	//过滤是否验证token
	//文档里我没给出utils.IsContainArr的代码，这个东西不重要，你直接删掉这段都行，这只是一个url过滤的逻辑
	/*if utils.IsContainArr(noVerify, c.Request.RequestURI) {
		return
	}*/
	//fmt.Println(noVerify, url)
	for _, v := range noVerify {
		if v == url {
			fmt.Println(url)
			//return true
		} else {
			fmt.Println("没有匹配项")
			//return false
		}
	}

	/*token := c.GetHeader("token")
	if token == "" {
		panic("token not exist !")
	}*/
	//验证token，并存储在请求中
	//c.Set("user", parseToken(token))
}
