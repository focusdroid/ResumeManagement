package test

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateTime(t *testing.T) {
	timer := time.Now()
	timera := time.Now().Unix()
	timed := time.Now().Add(time.Hour * 1).Unix()
	fmt.Println("===", timera, timed)
	fmt.Println(timera >= timed, timera < timed)
	createTime(int(timera))
	createTime(int(timed))
	fmt.Println("===")

	//temp := "2022-10-22 20:38:39.62498 +0800 CST m=+0.007803801"
	//fmt.Println(temp.Format("2006-01-02 15:04:05"))
	//timer1 := timer.Format("2006-01-02 15:04:05")
	//timer2 := timer.Format("2006-01-02 15:04:05")
	//fmt.Println(timer1)
	//fmt.Println(timer2)
	uninxtime := timer.Unix()
	fmt.Println(uninxtime)

	ut := 1666442453
	old := 1666443101
	news := old - ut
	fmt.Println(news / 1000)
	createTime(old)
	createTime(ut)

}

func createTime(ut int) {
	timer := time.Unix(int64(ut), 0).Format("2006-01-02 15:04:05")
	fmt.Println(timer)
}
