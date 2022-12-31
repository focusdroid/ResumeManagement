package test

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
	"testing"
	"time"
)

/*func RandomNumber(t *testing.T) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)
	return vcode
}*/

// 生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <focusdroid@163.com>"
	e.To = []string{"focusdroid@yeah.net"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "验证码发送测试"
	var num string = GetRand()
	fmt.Println(num)
	//e.Text = []byte("您的验证码是：<b>123456</b>")
	e.HTML = []byte("您的验证码是：<b>" + num + "</b>")
	//err := e.Send("smtp.163.com:465", smtp.PlainAuth("", "focusdroid@163.com", "MYBVJUDOLMJTSERW", "smtp.163.com"))
	// 返回EOF时候，关闭SSL重试, 授权需要获取授权码
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "focusdroid@163.com", "MYBVJUDOLMJTSERW", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
